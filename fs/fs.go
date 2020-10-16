// SPDX-License-Identifier: Apache-2.0
//
// Copyright 2020 EPAM Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fs

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// #cgo pkg-config: blkid
// #include <blkid.h>
import "C"

/*******************************************************************************
 * Consts
 ******************************************************************************/

const (
	retryCount = 3
	retryDelay = 1 * time.Second
)

const ioBufferSize = 1024 * 1024

const (
	copyBreathInterval = 5 * time.Second
	copyBreathTime     = 500 * time.Millisecond
)

const (
	tagTypeLabel    = "LABEL"
	tagTypeFSType   = "TYPE"
	tagTypePartUUID = "PARTUUID"
)

/*******************************************************************************
 * Types
 ******************************************************************************/

// PartInfo partition info
type PartInfo struct {
	PartUUID uuid.UUID
	FSType   string
	Label    string
}

/*******************************************************************************
 * Public
 ******************************************************************************/

// Mount creates mount point and mount device to it
func Mount(device string, mountPoint string, fsType string, flags uintptr, opts string) (err error) {
	log.WithFields(log.Fields{"device": device, "type": fsType, "mountPoint": mountPoint}).Debug("Mount partition")

	if err = os.MkdirAll(mountPoint, 0755); err != nil {
		return err
	}

	if err = retry(
		func() error {
			return syscall.Mount(device, mountPoint, fsType, flags, opts)
		},
		func(err error) {
			log.Warningf("Mount error: %s, try remount...", err)

			// Try to sync and force umount
			syscall.Unmount(mountPoint, syscall.MNT_FORCE)
		}); err != nil {
		return err
	}

	return nil
}

// Umount umount mount point and remove it
func Umount(mountPoint string) (err error) {
	log.WithFields(log.Fields{"mountPoint": mountPoint}).Debug("Umount partition")

	defer func() {
		if removeErr := os.RemoveAll(mountPoint); removeErr != nil {
			log.Errorf("Can't remove mount point: %s", removeErr)
			if err == nil {
				err = removeErr
			}
		}
	}()

	if err = retry(
		func() error {
			syscall.Sync()

			return syscall.Unmount(mountPoint, 0)
		},
		func(err error) {
			log.Warningf("Umount error: %s, retry...", err)

			time.Sleep(retryDelay)

			// Try to sync and force umount
			syscall.Sync()
		}); err != nil {

		log.Warningf("Can't umount for: %s", mountPoint)

		if output, err := exec.Command("lsof", mountPoint).CombinedOutput(); err == nil {
			log.Debugf("lsof says: %s", string(output))
		}

		if output, err := exec.Command("fuser", "-cu", mountPoint).CombinedOutput(); err == nil {
			log.Debugf("fuser says: %s", string(output))
		}

		return err
	}

	return nil
}

// Copy copies one partition to another
func Copy(dst, src string) (copied int64, err error) {
	log.WithFields(log.Fields{"src": src, "dst": dst}).Debug("Copy partition")

	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_RDWR, 0)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	var duration time.Duration

	if copied, duration, err = copyData(dstFile, srcFile); err != nil {
		return copied, err
	}

	log.WithFields(log.Fields{"copied": copied, "duration": duration}).Debug("Copy partition")

	return copied, nil
}

// CopyFromGzipArchive copies partition from archived file
func CopyFromGzipArchive(dst, src string) (copied int64, err error) {
	log.WithFields(log.Fields{"src": src, "dst": dst}).Debug("Copy partition from archive")

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_RDWR, 0)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	gz, err := gzip.NewReader(srcFile)
	if err != nil {
		return 0, err
	}
	defer gz.Close()

	var duration time.Duration

	if copied, duration, err = copyData(dstFile, gz); err != nil {
		return copied, err
	}

	log.WithFields(log.Fields{"copied": copied, "duration": duration}).Debug("Copy partition from archive")

	return copied, nil
}

// GetPartInfo returns partition info
func GetPartInfo(partDevice string) (partInfo PartInfo, err error) {
	var (
		blkdev   C.blkid_dev
		blkcache C.blkid_cache
	)

	if ret := C.blkid_get_cache(&blkcache, C.CString("/dev/null")); ret != 0 {
		return PartInfo{}, errors.New("can't get blkid cache")
	}

	if blkdev = C.blkid_get_dev(blkcache, C.CString(partDevice), C.BLKID_DEV_NORMAL); blkdev == nil {
		return PartInfo{}, errors.New("can't get blkid device")
	}

	iter := C.blkid_tag_iterate_begin(blkdev)

	var (
		tagType  *C.char
		tagValue *C.char
	)

	for C.blkid_tag_next(iter, &tagType, &tagValue) == 0 {
		switch C.GoString(tagType) {
		case tagTypeLabel:
			partInfo.Label = C.GoString(tagValue)

		case tagTypeFSType:
			partInfo.FSType = C.GoString(tagValue)

		case tagTypePartUUID:
			var err error

			if partInfo.PartUUID, err = uuid.Parse(C.GoString(tagValue)); err != nil {
				log.Errorf("Can't parse PARTUUID")
			}
		}
	}

	C.blkid_tag_iterate_end(iter)

	return partInfo, nil
}

// GetParentDevice returns partition parent device
// Example: input: /dev/nvme0n1p2 output: /dev/nvme0n1
func GetParentDevice(partitionPath string) (devPath string, err error) {
	partition, err := filepath.Rel("/dev", partitionPath)
	if err != nil {
		return "", err
	}

	items, err := ioutil.ReadDir("/sys/block")
	if err != nil {
		return "", err
	}

	for _, item := range items {
		subItems, err := ioutil.ReadDir(path.Join("/sys/block", item.Name()))
		if err != nil {
			return "", err
		}

		for _, subItem := range subItems {
			if subItem.Name() == partition {
				return path.Join("/dev", item.Name()), nil
			}
		}
	}

	return "", errors.New("can't determine parent device")
}

// GetPartitionNum returns partition number
// Example: input: /dev/nvme0n1p2 output: 2
func GetPartitionNum(partitionPath string) (num int, err error) {
	partition, err := filepath.Rel("/dev", partitionPath)
	if err != nil {
		return 0, err
	}

	sysPath := fmt.Sprintf("/sys/class/block/%s/partition", partition)
	//Check if file exists
	if _, err := os.Stat(sysPath); err != nil {
		return 0, err
	}

	b, err := ioutil.ReadFile(sysPath)
	if err != nil || len(b) == 0 {
		return 0, err
	}

	num, err = strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		return 0, err
	}

	return num, err
}

/*******************************************************************************
 * Private
 ******************************************************************************/

func retry(caller func() error, restorer func(error)) (err error) {
	i := 0

	for {
		if err = caller(); err == nil {
			return nil
		}

		if i >= retryCount {
			return err
		}

		if restorer != nil {
			restorer(err)
		}

		i++
	}
}

func copyData(dst io.Writer, src io.Reader) (copied int64, duration time.Duration, err error) {
	startTime := time.Now()
	buf := make([]byte, ioBufferSize)

	for err != io.EOF {
		var readCount int

		if readCount, err = src.Read(buf); err != nil && err != io.EOF {
			return copied, duration, err
		}

		if readCount > 0 {
			var writeCount int

			if writeCount, err = dst.Write(buf[:readCount]); err != nil {
				return copied, duration, err
			}

			copied = copied + int64(writeCount)
		}

		if time.Now().After(startTime.Add(duration).Add(copyBreathInterval)) {
			time.Sleep(copyBreathTime)

			duration = time.Since(startTime)

			log.WithFields(log.Fields{"copied": copied, "duration": duration}).Debug("Copy progress")
		}
	}

	duration = time.Since(startTime)

	return copied, duration, nil
}
