// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: provisioner/pb/provisioner.proto

package pb

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

type TerraformEvent int32

const (
	TerraformEvent_PLANNED_CHANGE TerraformEvent = 0
	TerraformEvent_CHANGE_SUMMARY TerraformEvent = 1
	TerraformEvent_APPLY_START    TerraformEvent = 2
	TerraformEvent_APPLY_PROGRESS TerraformEvent = 3
	TerraformEvent_APPLY_ERRORED  TerraformEvent = 4
	TerraformEvent_APPLY_COMPLETE TerraformEvent = 5
	TerraformEvent_DIAGNOSTIC     TerraformEvent = 6
)

// Enum value maps for TerraformEvent.
var (
	TerraformEvent_name = map[int32]string{
		0: "PLANNED_CHANGE",
		1: "CHANGE_SUMMARY",
		2: "APPLY_START",
		3: "APPLY_PROGRESS",
		4: "APPLY_ERRORED",
		5: "APPLY_COMPLETE",
		6: "DIAGNOSTIC",
	}
	TerraformEvent_value = map[string]int32{
		"PLANNED_CHANGE": 0,
		"CHANGE_SUMMARY": 1,
		"APPLY_START":    2,
		"APPLY_PROGRESS": 3,
		"APPLY_ERRORED":  4,
		"APPLY_COMPLETE": 5,
		"DIAGNOSTIC":     6,
	}
)

func (x TerraformEvent) Enum() *TerraformEvent {
	p := new(TerraformEvent)
	*p = x
	return p
}

func (x TerraformEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TerraformEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_provisioner_pb_provisioner_proto_enumTypes[0].Descriptor()
}

func (TerraformEvent) Type() protoreflect.EnumType {
	return &file_provisioner_pb_provisioner_proto_enumTypes[0]
}

func (x TerraformEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TerraformEvent.Descriptor instead.
func (TerraformEvent) EnumDescriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{0}
}

type TerraformStateMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TerraformStateMeta) Reset() {
	*x = TerraformStateMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_pb_provisioner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformStateMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformStateMeta) ProtoMessage() {}

func (x *TerraformStateMeta) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_pb_provisioner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformStateMeta.ProtoReflect.Descriptor instead.
func (*TerraformStateMeta) Descriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{0}
}

type Infra struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int64  `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Id        int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Suffix    string `protobuf:"bytes,3,opt,name=suffix,proto3" json:"suffix,omitempty"`
}

func (x *Infra) Reset() {
	*x = Infra{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_pb_provisioner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Infra) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Infra) ProtoMessage() {}

func (x *Infra) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_pb_provisioner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Infra.ProtoReflect.Descriptor instead.
func (*Infra) Descriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{1}
}

func (x *Infra) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *Infra) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Infra) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

type StateUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Update     string `protobuf:"bytes,2,opt,name=update,proto3" json:"update,omitempty"`
}

func (x *StateUpdate) Reset() {
	*x = StateUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_pb_provisioner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateUpdate) ProtoMessage() {}

func (x *StateUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_pb_provisioner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateUpdate.ProtoReflect.Descriptor instead.
func (*StateUpdate) Descriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{2}
}

func (x *StateUpdate) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *StateUpdate) GetUpdate() string {
	if x != nil {
		return x.Update
	}
	return ""
}

type TerraformResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr         string            `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Resource     string            `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	ResourceType string            `protobuf:"bytes,3,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceName string            `protobuf:"bytes,4,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	Provider     string            `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	Errored      *TerraformErrored `protobuf:"bytes,6,opt,name=errored,proto3" json:"errored,omitempty"`
}

func (x *TerraformResource) Reset() {
	*x = TerraformResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_pb_provisioner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformResource) ProtoMessage() {}

func (x *TerraformResource) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_pb_provisioner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformResource.ProtoReflect.Descriptor instead.
func (*TerraformResource) Descriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{3}
}

func (x *TerraformResource) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *TerraformResource) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *TerraformResource) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *TerraformResource) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *TerraformResource) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *TerraformResource) GetErrored() *TerraformErrored {
	if x != nil {
		return x.Errored
	}
	return nil
}

type TerraformErrored struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErroredOut   bool   `protobuf:"varint,1,opt,name=errored_out,json=erroredOut,proto3" json:"errored_out,omitempty"`
	ErrorSummary string `protobuf:"bytes,2,opt,name=error_summary,json=errorSummary,proto3" json:"error_summary,omitempty"`
}

func (x *TerraformErrored) Reset() {
	*x = TerraformErrored{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_pb_provisioner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformErrored) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformErrored) ProtoMessage() {}

func (x *TerraformErrored) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_pb_provisioner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformErrored.ProtoReflect.Descriptor instead.
func (*TerraformErrored) Descriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{4}
}

func (x *TerraformErrored) GetErroredOut() bool {
	if x != nil {
		return x.ErroredOut
	}
	return false
}

func (x *TerraformErrored) GetErrorSummary() string {
	if x != nil {
		return x.ErrorSummary
	}
	return ""
}

type TerraformHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *TerraformResource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *TerraformHook) Reset() {
	*x = TerraformHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_pb_provisioner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformHook) ProtoMessage() {}

func (x *TerraformHook) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_pb_provisioner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformHook.ProtoReflect.Descriptor instead.
func (*TerraformHook) Descriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{5}
}

func (x *TerraformHook) GetResource() *TerraformResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type TerraformChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *TerraformResource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Action   string             `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *TerraformChange) Reset() {
	*x = TerraformChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_pb_provisioner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformChange) ProtoMessage() {}

func (x *TerraformChange) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_pb_provisioner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformChange.ProtoReflect.Descriptor instead.
func (*TerraformChange) Descriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{6}
}

func (x *TerraformChange) GetResource() *TerraformResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *TerraformChange) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type TerraformChanges struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Add       int64  `protobuf:"varint,1,opt,name=add,proto3" json:"add,omitempty"`
	Change    int64  `protobuf:"varint,2,opt,name=change,proto3" json:"change,omitempty"`
	Remove    int64  `protobuf:"varint,3,opt,name=remove,proto3" json:"remove,omitempty"`
	Operation string `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *TerraformChanges) Reset() {
	*x = TerraformChanges{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_pb_provisioner_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformChanges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformChanges) ProtoMessage() {}

func (x *TerraformChanges) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_pb_provisioner_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformChanges.ProtoReflect.Descriptor instead.
func (*TerraformChanges) Descriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{7}
}

func (x *TerraformChanges) GetAdd() int64 {
	if x != nil {
		return x.Add
	}
	return 0
}

func (x *TerraformChanges) GetChange() int64 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *TerraformChanges) GetRemove() int64 {
	if x != nil {
		return x.Remove
	}
	return 0
}

func (x *TerraformChanges) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

type DiagnosticDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Severity string `protobuf:"bytes,1,opt,name=severity,proto3" json:"severity,omitempty"`
	Summary  string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DiagnosticDetail) Reset() {
	*x = DiagnosticDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_pb_provisioner_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiagnosticDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiagnosticDetail) ProtoMessage() {}

func (x *DiagnosticDetail) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_pb_provisioner_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiagnosticDetail.ProtoReflect.Descriptor instead.
func (*DiagnosticDetail) Descriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{8}
}

func (x *DiagnosticDetail) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *DiagnosticDetail) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *DiagnosticDetail) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type TerraformLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level      string            `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	Message    string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp  string            `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Type       TerraformEvent    `protobuf:"varint,4,opt,name=type,proto3,enum=TerraformEvent" json:"type,omitempty"`
	Hook       *TerraformHook    `protobuf:"bytes,5,opt,name=hook,proto3" json:"hook,omitempty"`
	Change     *TerraformChange  `protobuf:"bytes,6,opt,name=change,proto3" json:"change,omitempty"`
	Changes    *TerraformChanges `protobuf:"bytes,7,opt,name=changes,proto3" json:"changes,omitempty"`
	Diagnostic *DiagnosticDetail `protobuf:"bytes,8,opt,name=diagnostic,proto3" json:"diagnostic,omitempty"`
}

func (x *TerraformLog) Reset() {
	*x = TerraformLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_pb_provisioner_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformLog) ProtoMessage() {}

func (x *TerraformLog) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_pb_provisioner_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformLog.ProtoReflect.Descriptor instead.
func (*TerraformLog) Descriptor() ([]byte, []int) {
	return file_provisioner_pb_provisioner_proto_rawDescGZIP(), []int{9}
}

func (x *TerraformLog) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *TerraformLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TerraformLog) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *TerraformLog) GetType() TerraformEvent {
	if x != nil {
		return x.Type
	}
	return TerraformEvent_PLANNED_CHANGE
}

func (x *TerraformLog) GetHook() *TerraformHook {
	if x != nil {
		return x.Hook
	}
	return nil
}

func (x *TerraformLog) GetChange() *TerraformChange {
	if x != nil {
		return x.Change
	}
	return nil
}

func (x *TerraformLog) GetChanges() *TerraformChanges {
	if x != nil {
		return x.Changes
	}
	return nil
}

func (x *TerraformLog) GetDiagnostic() *DiagnosticDetail {
	if x != nil {
		return x.Diagnostic
	}
	return nil
}

var File_provisioner_pb_provisioner_proto protoreflect.FileDescriptor

var file_provisioner_pb_provisioner_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x72,
	0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0x46, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x22, 0xd6, 0x01, 0x0a, 0x11, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64,
	0x52, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x22, 0x58, 0x0a, 0x10, 0x54, 0x65, 0x72,
	0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x22, 0x3f, 0x0a, 0x0d, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x2e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x59, 0x0a, 0x0f, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72,
	0x6d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x54, 0x65, 0x72, 0x72,
	0x61, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x72, 0x0a, 0x10, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x61, 0x64, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x10, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xaf, 0x02, 0x0a, 0x0c, 0x54, 0x65, 0x72, 0x72,
	0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x68,
	0x6f, 0x6f, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x54, 0x65, 0x72, 0x72,
	0x61, 0x66, 0x6f, 0x72, 0x6d, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x68, 0x6f, 0x6f, 0x6b, 0x12,
	0x28, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x54, 0x65, 0x72,
	0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x44, 0x69, 0x61,
	0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0a, 0x64,
	0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2a, 0x94, 0x01, 0x0a, 0x0e, 0x54, 0x65,
	0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x0e,
	0x50, 0x4c, 0x41, 0x4e, 0x4e, 0x45, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x41,
	0x52, 0x59, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x50, 0x50,
	0x4c, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x45, 0x44, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e,
	0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05,
	0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x41, 0x47, 0x4e, 0x4f, 0x53, 0x54, 0x49, 0x43, 0x10, 0x06,
	0x32, 0x6d, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x1a, 0x0c, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x32, 0x0a, 0x08, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x0d, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x4c, 0x6f, 0x67, 0x1a, 0x13, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provisioner_pb_provisioner_proto_rawDescOnce sync.Once
	file_provisioner_pb_provisioner_proto_rawDescData = file_provisioner_pb_provisioner_proto_rawDesc
)

func file_provisioner_pb_provisioner_proto_rawDescGZIP() []byte {
	file_provisioner_pb_provisioner_proto_rawDescOnce.Do(func() {
		file_provisioner_pb_provisioner_proto_rawDescData = protoimpl.X.CompressGZIP(file_provisioner_pb_provisioner_proto_rawDescData)
	})
	return file_provisioner_pb_provisioner_proto_rawDescData
}

var file_provisioner_pb_provisioner_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_provisioner_pb_provisioner_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_provisioner_pb_provisioner_proto_goTypes = []interface{}{
	(TerraformEvent)(0),        // 0: TerraformEvent
	(*TerraformStateMeta)(nil), // 1: TerraformStateMeta
	(*Infra)(nil),              // 2: Infra
	(*StateUpdate)(nil),        // 3: StateUpdate
	(*TerraformResource)(nil),  // 4: TerraformResource
	(*TerraformErrored)(nil),   // 5: TerraformErrored
	(*TerraformHook)(nil),      // 6: TerraformHook
	(*TerraformChange)(nil),    // 7: TerraformChange
	(*TerraformChanges)(nil),   // 8: TerraformChanges
	(*DiagnosticDetail)(nil),   // 9: DiagnosticDetail
	(*TerraformLog)(nil),       // 10: TerraformLog
}
var file_provisioner_pb_provisioner_proto_depIdxs = []int32{
	5,  // 0: TerraformResource.errored:type_name -> TerraformErrored
	4,  // 1: TerraformHook.resource:type_name -> TerraformResource
	4,  // 2: TerraformChange.resource:type_name -> TerraformResource
	0,  // 3: TerraformLog.type:type_name -> TerraformEvent
	6,  // 4: TerraformLog.hook:type_name -> TerraformHook
	7,  // 5: TerraformLog.change:type_name -> TerraformChange
	8,  // 6: TerraformLog.changes:type_name -> TerraformChanges
	9,  // 7: TerraformLog.diagnostic:type_name -> DiagnosticDetail
	2,  // 8: Provisioner.GetStateUpdate:input_type -> Infra
	10, // 9: Provisioner.StoreLog:input_type -> TerraformLog
	3,  // 10: Provisioner.GetStateUpdate:output_type -> StateUpdate
	1,  // 11: Provisioner.StoreLog:output_type -> TerraformStateMeta
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_provisioner_pb_provisioner_proto_init() }
func file_provisioner_pb_provisioner_proto_init() {
	if File_provisioner_pb_provisioner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_provisioner_pb_provisioner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformStateMeta); i {
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
		file_provisioner_pb_provisioner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Infra); i {
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
		file_provisioner_pb_provisioner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateUpdate); i {
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
		file_provisioner_pb_provisioner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformResource); i {
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
		file_provisioner_pb_provisioner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformErrored); i {
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
		file_provisioner_pb_provisioner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformHook); i {
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
		file_provisioner_pb_provisioner_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformChange); i {
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
		file_provisioner_pb_provisioner_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformChanges); i {
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
		file_provisioner_pb_provisioner_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiagnosticDetail); i {
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
		file_provisioner_pb_provisioner_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformLog); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_provisioner_pb_provisioner_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provisioner_pb_provisioner_proto_goTypes,
		DependencyIndexes: file_provisioner_pb_provisioner_proto_depIdxs,
		EnumInfos:         file_provisioner_pb_provisioner_proto_enumTypes,
		MessageInfos:      file_provisioner_pb_provisioner_proto_msgTypes,
	}.Build()
	File_provisioner_pb_provisioner_proto = out.File
	file_provisioner_pb_provisioner_proto_rawDesc = nil
	file_provisioner_pb_provisioner_proto_goTypes = nil
	file_provisioner_pb_provisioner_proto_depIdxs = nil
}
