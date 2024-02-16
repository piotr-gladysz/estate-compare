// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: condition.proto

package api

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddConditionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Wasm []byte `protobuf:"bytes,2,opt,name=wasm,proto3" json:"wasm,omitempty"`
}

func (x *AddConditionRequest) Reset() {
	*x = AddConditionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_condition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddConditionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddConditionRequest) ProtoMessage() {}

func (x *AddConditionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_condition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddConditionRequest.ProtoReflect.Descriptor instead.
func (*AddConditionRequest) Descriptor() ([]byte, []int) {
	return file_condition_proto_rawDescGZIP(), []int{0}
}

func (x *AddConditionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddConditionRequest) GetWasm() []byte {
	if x != nil {
		return x.Wasm
	}
	return nil
}

type GetConditionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetConditionRequest) Reset() {
	*x = GetConditionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_condition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConditionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConditionRequest) ProtoMessage() {}

func (x *GetConditionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_condition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConditionRequest.ProtoReflect.Descriptor instead.
func (*GetConditionRequest) Descriptor() ([]byte, []int) {
	return file_condition_proto_rawDescGZIP(), []int{1}
}

func (x *GetConditionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetConditionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetConditionsRequest) Reset() {
	*x = GetConditionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_condition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConditionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConditionsRequest) ProtoMessage() {}

func (x *GetConditionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_condition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConditionsRequest.ProtoReflect.Descriptor instead.
func (*GetConditionsRequest) Descriptor() ([]byte, []int) {
	return file_condition_proto_rawDescGZIP(), []int{2}
}

func (x *GetConditionsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetConditionsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ConditionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Created int64  `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated int64  `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *ConditionResponse) Reset() {
	*x = ConditionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_condition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConditionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConditionResponse) ProtoMessage() {}

func (x *ConditionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_condition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConditionResponse.ProtoReflect.Descriptor instead.
func (*ConditionResponse) Descriptor() ([]byte, []int) {
	return file_condition_proto_rawDescGZIP(), []int{3}
}

func (x *ConditionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConditionResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConditionResponse) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *ConditionResponse) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type ConditionsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conditions []*ConditionResponse `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
	Total      int32                `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ConditionsListResponse) Reset() {
	*x = ConditionsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_condition_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConditionsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConditionsListResponse) ProtoMessage() {}

func (x *ConditionsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_condition_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConditionsListResponse.ProtoReflect.Descriptor instead.
func (*ConditionsListResponse) Descriptor() ([]byte, []int) {
	return file_condition_proto_rawDescGZIP(), []int{4}
}

func (x *ConditionsListResponse) GetConditions() []*ConditionResponse {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *ConditionsListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_condition_proto protoreflect.FileDescriptor

var file_condition_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x77, 0x61, 0x73, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x77,
	0x61, 0x73, 0x6d, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02,
	0x28, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04,
	0x1a, 0x02, 0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6b,
	0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x66, 0x0a, 0x16, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x32, 0xdf, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x64, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x72, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x42, 0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x69, 0x6f, 0x74, 0x72, 0x2d, 0x67, 0x6c, 0x61, 0x64, 0x79, 0x73, 0x7a, 0x2f, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0xa2,
	0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x03, 0x41, 0x70, 0x69, 0xca, 0x02, 0x03, 0x41, 0x70,
	0x69, 0xe2, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x03, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_condition_proto_rawDescOnce sync.Once
	file_condition_proto_rawDescData = file_condition_proto_rawDesc
)

func file_condition_proto_rawDescGZIP() []byte {
	file_condition_proto_rawDescOnce.Do(func() {
		file_condition_proto_rawDescData = protoimpl.X.CompressGZIP(file_condition_proto_rawDescData)
	})
	return file_condition_proto_rawDescData
}

var file_condition_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_condition_proto_goTypes = []interface{}{
	(*AddConditionRequest)(nil),    // 0: api.AddConditionRequest
	(*GetConditionRequest)(nil),    // 1: api.GetConditionRequest
	(*GetConditionsRequest)(nil),   // 2: api.GetConditionsRequest
	(*ConditionResponse)(nil),      // 3: api.ConditionResponse
	(*ConditionsListResponse)(nil), // 4: api.ConditionsListResponse
}
var file_condition_proto_depIdxs = []int32{
	3, // 0: api.ConditionsListResponse.conditions:type_name -> api.ConditionResponse
	0, // 1: api.ConditionService.AddCondition:input_type -> api.AddConditionRequest
	1, // 2: api.ConditionService.GetCondition:input_type -> api.GetConditionRequest
	2, // 3: api.ConditionService.GetConditions:input_type -> api.GetConditionsRequest
	3, // 4: api.ConditionService.AddCondition:output_type -> api.ConditionResponse
	3, // 5: api.ConditionService.GetCondition:output_type -> api.ConditionResponse
	4, // 6: api.ConditionService.GetConditions:output_type -> api.ConditionsListResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_condition_proto_init() }
func file_condition_proto_init() {
	if File_condition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_condition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddConditionRequest); i {
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
		file_condition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConditionRequest); i {
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
		file_condition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConditionsRequest); i {
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
		file_condition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionResponse); i {
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
		file_condition_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionsListResponse); i {
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
			RawDescriptor: file_condition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_condition_proto_goTypes,
		DependencyIndexes: file_condition_proto_depIdxs,
		MessageInfos:      file_condition_proto_msgTypes,
	}.Build()
	File_condition_proto = out.File
	file_condition_proto_rawDesc = nil
	file_condition_proto_goTypes = nil
	file_condition_proto_depIdxs = nil
}
