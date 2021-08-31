// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: updatemanager/updatemanager.proto

package updatemanager

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UmState int32

const (
	UmState_IDLE     UmState = 0
	UmState_PREPARED UmState = 1
	UmState_UPDATED  UmState = 2
	UmState_FAILED   UmState = 3
)

// Enum value maps for UmState.
var (
	UmState_name = map[int32]string{
		0: "IDLE",
		1: "PREPARED",
		2: "UPDATED",
		3: "FAILED",
	}
	UmState_value = map[string]int32{
		"IDLE":     0,
		"PREPARED": 1,
		"UPDATED":  2,
		"FAILED":   3,
	}
)

func (x UmState) Enum() *UmState {
	p := new(UmState)
	*p = x
	return p
}

func (x UmState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UmState) Descriptor() protoreflect.EnumDescriptor {
	return file_updatemanager_updatemanager_proto_enumTypes[0].Descriptor()
}

func (UmState) Type() protoreflect.EnumType {
	return &file_updatemanager_updatemanager_proto_enumTypes[0]
}

func (x UmState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UmState.Descriptor instead.
func (UmState) EnumDescriptor() ([]byte, []int) {
	return file_updatemanager_updatemanager_proto_rawDescGZIP(), []int{0}
}

type ComponentStatus int32

const (
	ComponentStatus_INSTALLED  ComponentStatus = 0
	ComponentStatus_INSTALLING ComponentStatus = 1
	ComponentStatus_ERROR      ComponentStatus = 2
)

// Enum value maps for ComponentStatus.
var (
	ComponentStatus_name = map[int32]string{
		0: "INSTALLED",
		1: "INSTALLING",
		2: "ERROR",
	}
	ComponentStatus_value = map[string]int32{
		"INSTALLED":  0,
		"INSTALLING": 1,
		"ERROR":      2,
	}
)

func (x ComponentStatus) Enum() *ComponentStatus {
	p := new(ComponentStatus)
	*p = x
	return p
}

func (x ComponentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComponentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_updatemanager_updatemanager_proto_enumTypes[1].Descriptor()
}

func (ComponentStatus) Type() protoreflect.EnumType {
	return &file_updatemanager_updatemanager_proto_enumTypes[1]
}

func (x ComponentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComponentStatus.Descriptor instead.
func (ComponentStatus) EnumDescriptor() ([]byte, []int) {
	return file_updatemanager_updatemanager_proto_rawDescGZIP(), []int{1}
}

type CMMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to CMMessage:
	//	*CMMessages_PrepareUpdate
	//	*CMMessages_StartUpdate
	//	*CMMessages_ApplyUpdate
	//	*CMMessages_RevertUpdate
	CMMessage isCMMessages_CMMessage `protobuf_oneof:"CMMessage"`
}

func (x *CMMessages) Reset() {
	*x = CMMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updatemanager_updatemanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMMessages) ProtoMessage() {}

func (x *CMMessages) ProtoReflect() protoreflect.Message {
	mi := &file_updatemanager_updatemanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMMessages.ProtoReflect.Descriptor instead.
func (*CMMessages) Descriptor() ([]byte, []int) {
	return file_updatemanager_updatemanager_proto_rawDescGZIP(), []int{0}
}

func (m *CMMessages) GetCMMessage() isCMMessages_CMMessage {
	if m != nil {
		return m.CMMessage
	}
	return nil
}

func (x *CMMessages) GetPrepareUpdate() *PrepareUpdate {
	if x, ok := x.GetCMMessage().(*CMMessages_PrepareUpdate); ok {
		return x.PrepareUpdate
	}
	return nil
}

func (x *CMMessages) GetStartUpdate() *StartUpdate {
	if x, ok := x.GetCMMessage().(*CMMessages_StartUpdate); ok {
		return x.StartUpdate
	}
	return nil
}

func (x *CMMessages) GetApplyUpdate() *ApplyUpdate {
	if x, ok := x.GetCMMessage().(*CMMessages_ApplyUpdate); ok {
		return x.ApplyUpdate
	}
	return nil
}

func (x *CMMessages) GetRevertUpdate() *RevertUpdate {
	if x, ok := x.GetCMMessage().(*CMMessages_RevertUpdate); ok {
		return x.RevertUpdate
	}
	return nil
}

type isCMMessages_CMMessage interface {
	isCMMessages_CMMessage()
}

type CMMessages_PrepareUpdate struct {
	PrepareUpdate *PrepareUpdate `protobuf:"bytes,1,opt,name=prepare_update,json=prepareUpdate,proto3,oneof"`
}

type CMMessages_StartUpdate struct {
	StartUpdate *StartUpdate `protobuf:"bytes,2,opt,name=start_update,json=startUpdate,proto3,oneof"`
}

type CMMessages_ApplyUpdate struct {
	ApplyUpdate *ApplyUpdate `protobuf:"bytes,3,opt,name=apply_update,json=applyUpdate,proto3,oneof"`
}

type CMMessages_RevertUpdate struct {
	RevertUpdate *RevertUpdate `protobuf:"bytes,4,opt,name=revert_update,json=revertUpdate,proto3,oneof"`
}

func (*CMMessages_PrepareUpdate) isCMMessages_CMMessage() {}

func (*CMMessages_StartUpdate) isCMMessages_CMMessage() {}

func (*CMMessages_ApplyUpdate) isCMMessages_CMMessage() {}

func (*CMMessages_RevertUpdate) isCMMessages_CMMessage() {}

type PrepareComponentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VendorVersion string `protobuf:"bytes,2,opt,name=vendor_version,json=vendorVersion,proto3" json:"vendor_version,omitempty"`
	AosVersion    uint64 `protobuf:"varint,3,opt,name=aos_version,json=aosVersion,proto3" json:"aos_version,omitempty"`
	Annotations   string `protobuf:"bytes,4,opt,name=annotations,proto3" json:"annotations,omitempty"`
	Url           string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Sha256        []byte `protobuf:"bytes,6,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Sha512        []byte `protobuf:"bytes,7,opt,name=sha512,proto3" json:"sha512,omitempty"`
	Size          uint64 `protobuf:"varint,8,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *PrepareComponentInfo) Reset() {
	*x = PrepareComponentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updatemanager_updatemanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareComponentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareComponentInfo) ProtoMessage() {}

func (x *PrepareComponentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_updatemanager_updatemanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareComponentInfo.ProtoReflect.Descriptor instead.
func (*PrepareComponentInfo) Descriptor() ([]byte, []int) {
	return file_updatemanager_updatemanager_proto_rawDescGZIP(), []int{1}
}

func (x *PrepareComponentInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PrepareComponentInfo) GetVendorVersion() string {
	if x != nil {
		return x.VendorVersion
	}
	return ""
}

func (x *PrepareComponentInfo) GetAosVersion() uint64 {
	if x != nil {
		return x.AosVersion
	}
	return 0
}

func (x *PrepareComponentInfo) GetAnnotations() string {
	if x != nil {
		return x.Annotations
	}
	return ""
}

func (x *PrepareComponentInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PrepareComponentInfo) GetSha256() []byte {
	if x != nil {
		return x.Sha256
	}
	return nil
}

func (x *PrepareComponentInfo) GetSha512() []byte {
	if x != nil {
		return x.Sha512
	}
	return nil
}

func (x *PrepareComponentInfo) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type PrepareUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Components []*PrepareComponentInfo `protobuf:"bytes,1,rep,name=components,proto3" json:"components,omitempty"`
}

func (x *PrepareUpdate) Reset() {
	*x = PrepareUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updatemanager_updatemanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareUpdate) ProtoMessage() {}

func (x *PrepareUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_updatemanager_updatemanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareUpdate.ProtoReflect.Descriptor instead.
func (*PrepareUpdate) Descriptor() ([]byte, []int) {
	return file_updatemanager_updatemanager_proto_rawDescGZIP(), []int{2}
}

func (x *PrepareUpdate) GetComponents() []*PrepareComponentInfo {
	if x != nil {
		return x.Components
	}
	return nil
}

type StartUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartUpdate) Reset() {
	*x = StartUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updatemanager_updatemanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartUpdate) ProtoMessage() {}

func (x *StartUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_updatemanager_updatemanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartUpdate.ProtoReflect.Descriptor instead.
func (*StartUpdate) Descriptor() ([]byte, []int) {
	return file_updatemanager_updatemanager_proto_rawDescGZIP(), []int{3}
}

type ApplyUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApplyUpdate) Reset() {
	*x = ApplyUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updatemanager_updatemanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyUpdate) ProtoMessage() {}

func (x *ApplyUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_updatemanager_updatemanager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyUpdate.ProtoReflect.Descriptor instead.
func (*ApplyUpdate) Descriptor() ([]byte, []int) {
	return file_updatemanager_updatemanager_proto_rawDescGZIP(), []int{4}
}

type RevertUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RevertUpdate) Reset() {
	*x = RevertUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updatemanager_updatemanager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevertUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevertUpdate) ProtoMessage() {}

func (x *RevertUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_updatemanager_updatemanager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevertUpdate.ProtoReflect.Descriptor instead.
func (*RevertUpdate) Descriptor() ([]byte, []int) {
	return file_updatemanager_updatemanager_proto_rawDescGZIP(), []int{5}
}

type SystemComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VendorVersion string          `protobuf:"bytes,2,opt,name=vendor_version,json=vendorVersion,proto3" json:"vendor_version,omitempty"`
	AosVersion    uint64          `protobuf:"varint,3,opt,name=aos_version,json=aosVersion,proto3" json:"aos_version,omitempty"`
	Status        ComponentStatus `protobuf:"varint,4,opt,name=status,proto3,enum=updatemanager.ComponentStatus" json:"status,omitempty"`
	Error         string          `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SystemComponent) Reset() {
	*x = SystemComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updatemanager_updatemanager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemComponent) ProtoMessage() {}

func (x *SystemComponent) ProtoReflect() protoreflect.Message {
	mi := &file_updatemanager_updatemanager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemComponent.ProtoReflect.Descriptor instead.
func (*SystemComponent) Descriptor() ([]byte, []int) {
	return file_updatemanager_updatemanager_proto_rawDescGZIP(), []int{6}
}

func (x *SystemComponent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SystemComponent) GetVendorVersion() string {
	if x != nil {
		return x.VendorVersion
	}
	return ""
}

func (x *SystemComponent) GetAosVersion() uint64 {
	if x != nil {
		return x.AosVersion
	}
	return 0
}

func (x *SystemComponent) GetStatus() ComponentStatus {
	if x != nil {
		return x.Status
	}
	return ComponentStatus_INSTALLED
}

func (x *SystemComponent) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UmId       string             `protobuf:"bytes,1,opt,name=um_id,json=umId,proto3" json:"um_id,omitempty"`
	UmState    UmState            `protobuf:"varint,2,opt,name=um_state,json=umState,proto3,enum=updatemanager.UmState" json:"um_state,omitempty"`
	Error      string             `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Components []*SystemComponent `protobuf:"bytes,4,rep,name=components,proto3" json:"components,omitempty"`
}

func (x *UpdateStatus) Reset() {
	*x = UpdateStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updatemanager_updatemanager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatus) ProtoMessage() {}

func (x *UpdateStatus) ProtoReflect() protoreflect.Message {
	mi := &file_updatemanager_updatemanager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatus.ProtoReflect.Descriptor instead.
func (*UpdateStatus) Descriptor() ([]byte, []int) {
	return file_updatemanager_updatemanager_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateStatus) GetUmId() string {
	if x != nil {
		return x.UmId
	}
	return ""
}

func (x *UpdateStatus) GetUmState() UmState {
	if x != nil {
		return x.UmState
	}
	return UmState_IDLE
}

func (x *UpdateStatus) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *UpdateStatus) GetComponents() []*SystemComponent {
	if x != nil {
		return x.Components
	}
	return nil
}

var File_updatemanager_updatemanager_proto protoreflect.FileDescriptor

var file_updatemanager_updatemanager_proto_rawDesc = []byte{
	0x0a, 0x21, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x22, 0xa6, 0x02, 0x0a, 0x0a, 0x43, 0x4d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x45, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00,
	0x52, 0x0c, 0x72, 0x65, 0x76, 0x65, 0x72, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x43, 0x4d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe6, 0x01, 0x0a, 0x14,
	0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x61, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x35,
	0x31, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x68, 0x61, 0x35, 0x31, 0x32,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x54, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x65, 0x76, 0x65,
	0x72, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x0f, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x6f, 0x73, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0xac, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x75, 0x6d, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x07, 0x75, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x2a, 0x3a, 0x0a, 0x07, 0x55, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x44, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x3b, 0x0a,
	0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x32, 0x5a, 0x0a, 0x0c, 0x55, 0x4d,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0a, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x4d, 0x12, 0x1b, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x19, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x4d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x70, 0x63, 0x74,
	0x2e, 0x65, 0x70, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x70, 0x6d, 0x64, 0x2d, 0x61,
	0x65, 0x70, 0x72, 0x2f, 0x61, 0x6f, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_updatemanager_updatemanager_proto_rawDescOnce sync.Once
	file_updatemanager_updatemanager_proto_rawDescData = file_updatemanager_updatemanager_proto_rawDesc
)

func file_updatemanager_updatemanager_proto_rawDescGZIP() []byte {
	file_updatemanager_updatemanager_proto_rawDescOnce.Do(func() {
		file_updatemanager_updatemanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_updatemanager_updatemanager_proto_rawDescData)
	})
	return file_updatemanager_updatemanager_proto_rawDescData
}

var file_updatemanager_updatemanager_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_updatemanager_updatemanager_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_updatemanager_updatemanager_proto_goTypes = []interface{}{
	(UmState)(0),                 // 0: updatemanager.UmState
	(ComponentStatus)(0),         // 1: updatemanager.ComponentStatus
	(*CMMessages)(nil),           // 2: updatemanager.CMMessages
	(*PrepareComponentInfo)(nil), // 3: updatemanager.PrepareComponentInfo
	(*PrepareUpdate)(nil),        // 4: updatemanager.PrepareUpdate
	(*StartUpdate)(nil),          // 5: updatemanager.StartUpdate
	(*ApplyUpdate)(nil),          // 6: updatemanager.ApplyUpdate
	(*RevertUpdate)(nil),         // 7: updatemanager.RevertUpdate
	(*SystemComponent)(nil),      // 8: updatemanager.SystemComponent
	(*UpdateStatus)(nil),         // 9: updatemanager.UpdateStatus
}
var file_updatemanager_updatemanager_proto_depIdxs = []int32{
	4, // 0: updatemanager.CMMessages.prepare_update:type_name -> updatemanager.PrepareUpdate
	5, // 1: updatemanager.CMMessages.start_update:type_name -> updatemanager.StartUpdate
	6, // 2: updatemanager.CMMessages.apply_update:type_name -> updatemanager.ApplyUpdate
	7, // 3: updatemanager.CMMessages.revert_update:type_name -> updatemanager.RevertUpdate
	3, // 4: updatemanager.PrepareUpdate.components:type_name -> updatemanager.PrepareComponentInfo
	1, // 5: updatemanager.SystemComponent.status:type_name -> updatemanager.ComponentStatus
	0, // 6: updatemanager.UpdateStatus.um_state:type_name -> updatemanager.UmState
	8, // 7: updatemanager.UpdateStatus.components:type_name -> updatemanager.SystemComponent
	9, // 8: updatemanager.UMController.RegisterUM:input_type -> updatemanager.UpdateStatus
	2, // 9: updatemanager.UMController.RegisterUM:output_type -> updatemanager.CMMessages
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_updatemanager_updatemanager_proto_init() }
func file_updatemanager_updatemanager_proto_init() {
	if File_updatemanager_updatemanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_updatemanager_updatemanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMMessages); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_updatemanager_updatemanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareComponentInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_updatemanager_updatemanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_updatemanager_updatemanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_updatemanager_updatemanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_updatemanager_updatemanager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevertUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_updatemanager_updatemanager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemComponent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_updatemanager_updatemanager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_updatemanager_updatemanager_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CMMessages_PrepareUpdate)(nil),
		(*CMMessages_StartUpdate)(nil),
		(*CMMessages_ApplyUpdate)(nil),
		(*CMMessages_RevertUpdate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_updatemanager_updatemanager_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_updatemanager_updatemanager_proto_goTypes,
		DependencyIndexes: file_updatemanager_updatemanager_proto_depIdxs,
		EnumInfos:         file_updatemanager_updatemanager_proto_enumTypes,
		MessageInfos:      file_updatemanager_updatemanager_proto_msgTypes,
	}.Build()
	File_updatemanager_updatemanager_proto = out.File
	file_updatemanager_updatemanager_proto_rawDesc = nil
	file_updatemanager_updatemanager_proto_goTypes = nil
	file_updatemanager_updatemanager_proto_depIdxs = nil
}
