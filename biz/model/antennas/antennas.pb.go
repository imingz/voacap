// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.26.1
// source: antennas.proto

package antennas

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	_ "voacap/biz/model/api"
	common "voacap/biz/model/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAntennasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAntennasRequest) Reset() {
	*x = GetAntennasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antennas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAntennasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAntennasRequest) ProtoMessage() {}

func (x *GetAntennasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antennas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAntennasRequest.ProtoReflect.Descriptor instead.
func (*GetAntennasRequest) Descriptor() ([]byte, []int) {
	return file_antennas_proto_rawDescGZIP(), []int{0}
}

type GetAntennasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" form:"code" query:"code"`
	Msg  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty" form:"msg" query:"msg"`
	Data []*common.Antenna `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" form:"data" query:"data"`
}

func (x *GetAntennasResponse) Reset() {
	*x = GetAntennasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antennas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAntennasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAntennasResponse) ProtoMessage() {}

func (x *GetAntennasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_antennas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAntennasResponse.ProtoReflect.Descriptor instead.
func (*GetAntennasResponse) Descriptor() ([]byte, []int) {
	return file_antennas_proto_rawDescGZIP(), []int{1}
}

func (x *GetAntennasResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAntennasResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetAntennasResponse) GetData() []*common.Antenna {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddAntennaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aname     string `protobuf:"bytes,2,opt,name=Aname,proto3" json:"Aname,omitempty" form:"Aname" query:"Aname"`
	Afile     string `protobuf:"bytes,3,opt,name=Afile,proto3" json:"Afile,omitempty" form:"Afile" query:"Afile"`
	AfbandMin int64  `protobuf:"varint,4,opt,name=AfbandMin,proto3" json:"AfbandMin,omitempty" form:"AfbandMin" query:"AfbandMin"`
	AfbandMax int64  `protobuf:"varint,5,opt,name=AfbandMax,proto3" json:"AfbandMax,omitempty" form:"AfbandMax" query:"AfbandMax"`
}

func (x *AddAntennaRequest) Reset() {
	*x = AddAntennaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antennas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAntennaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAntennaRequest) ProtoMessage() {}

func (x *AddAntennaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antennas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAntennaRequest.ProtoReflect.Descriptor instead.
func (*AddAntennaRequest) Descriptor() ([]byte, []int) {
	return file_antennas_proto_rawDescGZIP(), []int{2}
}

func (x *AddAntennaRequest) GetAname() string {
	if x != nil {
		return x.Aname
	}
	return ""
}

func (x *AddAntennaRequest) GetAfile() string {
	if x != nil {
		return x.Afile
	}
	return ""
}

func (x *AddAntennaRequest) GetAfbandMin() int64 {
	if x != nil {
		return x.AfbandMin
	}
	return 0
}

func (x *AddAntennaRequest) GetAfbandMax() int64 {
	if x != nil {
		return x.AfbandMax
	}
	return 0
}

type AddAntennaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" form:"code" query:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty" form:"msg" query:"msg"`
}

func (x *AddAntennaResponse) Reset() {
	*x = AddAntennaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antennas_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAntennaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAntennaResponse) ProtoMessage() {}

func (x *AddAntennaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_antennas_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAntennaResponse.ProtoReflect.Descriptor instead.
func (*AddAntennaResponse) Descriptor() ([]byte, []int) {
	return file_antennas_proto_rawDescGZIP(), []int{3}
}

func (x *AddAntennaResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AddAntennaResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UpdateAntennaByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AntennaId int64  `protobuf:"varint,1,opt,name=antennaId,proto3" json:"antennaId,omitempty" form:"antennaId" query:"antennaId"`
	Aname     string `protobuf:"bytes,2,opt,name=Aname,proto3" json:"Aname,omitempty" form:"Aname" query:"Aname"`
	Afile     string `protobuf:"bytes,3,opt,name=Afile,proto3" json:"Afile,omitempty" form:"Afile" query:"Afile"`
	AfbandMin int64  `protobuf:"varint,4,opt,name=AfbandMin,proto3" json:"AfbandMin,omitempty" form:"AfbandMin" query:"AfbandMin"`
	AfbandMax int64  `protobuf:"varint,5,opt,name=AfbandMax,proto3" json:"AfbandMax,omitempty" form:"AfbandMax" query:"AfbandMax"`
}

func (x *UpdateAntennaByIdRequest) Reset() {
	*x = UpdateAntennaByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antennas_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAntennaByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAntennaByIdRequest) ProtoMessage() {}

func (x *UpdateAntennaByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antennas_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAntennaByIdRequest.ProtoReflect.Descriptor instead.
func (*UpdateAntennaByIdRequest) Descriptor() ([]byte, []int) {
	return file_antennas_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAntennaByIdRequest) GetAntennaId() int64 {
	if x != nil {
		return x.AntennaId
	}
	return 0
}

func (x *UpdateAntennaByIdRequest) GetAname() string {
	if x != nil {
		return x.Aname
	}
	return ""
}

func (x *UpdateAntennaByIdRequest) GetAfile() string {
	if x != nil {
		return x.Afile
	}
	return ""
}

func (x *UpdateAntennaByIdRequest) GetAfbandMin() int64 {
	if x != nil {
		return x.AfbandMin
	}
	return 0
}

func (x *UpdateAntennaByIdRequest) GetAfbandMax() int64 {
	if x != nil {
		return x.AfbandMax
	}
	return 0
}

type UpdateAntennaByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" form:"code" query:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty" form:"msg" query:"msg"`
}

func (x *UpdateAntennaByIdResponse) Reset() {
	*x = UpdateAntennaByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antennas_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAntennaByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAntennaByIdResponse) ProtoMessage() {}

func (x *UpdateAntennaByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_antennas_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAntennaByIdResponse.ProtoReflect.Descriptor instead.
func (*UpdateAntennaByIdResponse) Descriptor() ([]byte, []int) {
	return file_antennas_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAntennaByIdResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateAntennaByIdResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DeleteAntennaByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AntennaId int64 `protobuf:"varint,1,opt,name=antennaId,proto3" json:"antennaId,omitempty" form:"antennaId" query:"antennaId"`
}

func (x *DeleteAntennaByIdRequest) Reset() {
	*x = DeleteAntennaByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antennas_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAntennaByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAntennaByIdRequest) ProtoMessage() {}

func (x *DeleteAntennaByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antennas_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAntennaByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteAntennaByIdRequest) Descriptor() ([]byte, []int) {
	return file_antennas_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAntennaByIdRequest) GetAntennaId() int64 {
	if x != nil {
		return x.AntennaId
	}
	return 0
}

type DeleteAntennaByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" form:"code" query:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty" form:"msg" query:"msg"`
}

func (x *DeleteAntennaByIdResponse) Reset() {
	*x = DeleteAntennaByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antennas_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAntennaByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAntennaByIdResponse) ProtoMessage() {}

func (x *DeleteAntennaByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_antennas_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAntennaByIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteAntennaByIdResponse) Descriptor() ([]byte, []int) {
	return file_antennas_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAntennaByIdResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteAntennaByIdResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_antennas_proto protoreflect.FileDescriptor

var file_antennas_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e,
	0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x60, 0x0a, 0x13, 0x67, 0x65, 0x74,
	0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x6e,
	0x74, 0x65, 0x6e, 0x6e, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7b, 0x0a, 0x11, 0x61,
	0x64, 0x64, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x41, 0x66, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x41, 0x66, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x66,
	0x62, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41,
	0x66, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x22, 0x3a, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x41,
	0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0xa0, 0x01, 0x0a, 0x18, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x66, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x41, 0x66, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x66, 0x62,
	0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x66,
	0x62, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x22, 0x41, 0x0a, 0x19, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x38, 0x0a, 0x18, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e,
	0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6e, 0x74, 0x65, 0x6e,
	0x6e, 0x61, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x19, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6e,
	0x74, 0x65, 0x6e, 0x6e, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xd9, 0x03, 0x0a, 0x0f, 0x41, 0x6e, 0x74, 0x65,
	0x6e, 0x6e, 0x61, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x6e, 0x74,
	0x65, 0x6e, 0x6e, 0x61, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x6e, 0x74, 0x65, 0x6e,
	0x6e, 0x61, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0xca, 0xc1, 0x18, 0x15, 0x2f, 0x61, 0x6e,
	0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e,
	0x61, 0x73, 0x12, 0x61, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61,
	0x12, 0x1b, 0x2e, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x41,
	0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x41, 0x6e, 0x74, 0x65,
	0x6e, 0x6e, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0xd2, 0xc1, 0x18,
	0x14, 0x2f, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x41, 0x6e,
	0x74, 0x65, 0x6e, 0x6e, 0x61, 0x12, 0x7d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x42, 0x79, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x61, 0x6e, 0x74,
	0x65, 0x6e, 0x6e, 0x61, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x74, 0x65,
	0x6e, 0x6e, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1f, 0xd2, 0xc1, 0x18, 0x1b, 0x2f, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e,
	0x61, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x7d, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6e,
	0x74, 0x65, 0x6e, 0x6e, 0x61, 0x42, 0x79, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x61, 0x6e, 0x74, 0x65,
	0x6e, 0x6e, 0x61, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6e, 0x74, 0x65, 0x6e,
	0x6e, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0xd2, 0xc1, 0x18, 0x1b, 0x2f, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61,
	0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x42,
	0x79, 0x49, 0x64, 0x42, 0x1b, 0x5a, 0x19, 0x76, 0x6f, 0x61, 0x63, 0x61, 0x70, 0x2f, 0x62, 0x69,
	0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_antennas_proto_rawDescOnce sync.Once
	file_antennas_proto_rawDescData = file_antennas_proto_rawDesc
)

func file_antennas_proto_rawDescGZIP() []byte {
	file_antennas_proto_rawDescOnce.Do(func() {
		file_antennas_proto_rawDescData = protoimpl.X.CompressGZIP(file_antennas_proto_rawDescData)
	})
	return file_antennas_proto_rawDescData
}

var file_antennas_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_antennas_proto_goTypes = []interface{}{
	(*GetAntennasRequest)(nil),        // 0: antennas.getAntennasRequest
	(*GetAntennasResponse)(nil),       // 1: antennas.getAntennasResponse
	(*AddAntennaRequest)(nil),         // 2: antennas.addAntennaRequest
	(*AddAntennaResponse)(nil),        // 3: antennas.addAntennaResponse
	(*UpdateAntennaByIdRequest)(nil),  // 4: antennas.updateAntennaByIdRequest
	(*UpdateAntennaByIdResponse)(nil), // 5: antennas.updateAntennaByIdResponse
	(*DeleteAntennaByIdRequest)(nil),  // 6: antennas.deleteAntennaByIdRequest
	(*DeleteAntennaByIdResponse)(nil), // 7: antennas.deleteAntennaByIdResponse
	(*common.Antenna)(nil),            // 8: common.Antenna
}
var file_antennas_proto_depIdxs = []int32{
	8, // 0: antennas.getAntennasResponse.data:type_name -> common.Antenna
	0, // 1: antennas.AntennasService.GetAntennas:input_type -> antennas.getAntennasRequest
	2, // 2: antennas.AntennasService.AddAntenna:input_type -> antennas.addAntennaRequest
	4, // 3: antennas.AntennasService.UpdateAntennaById:input_type -> antennas.updateAntennaByIdRequest
	6, // 4: antennas.AntennasService.DeleteAntennaById:input_type -> antennas.deleteAntennaByIdRequest
	1, // 5: antennas.AntennasService.GetAntennas:output_type -> antennas.getAntennasResponse
	3, // 6: antennas.AntennasService.AddAntenna:output_type -> antennas.addAntennaResponse
	5, // 7: antennas.AntennasService.UpdateAntennaById:output_type -> antennas.updateAntennaByIdResponse
	7, // 8: antennas.AntennasService.DeleteAntennaById:output_type -> antennas.deleteAntennaByIdResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_antennas_proto_init() }
func file_antennas_proto_init() {
	if File_antennas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_antennas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAntennasRequest); i {
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
		file_antennas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAntennasResponse); i {
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
		file_antennas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAntennaRequest); i {
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
		file_antennas_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAntennaResponse); i {
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
		file_antennas_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAntennaByIdRequest); i {
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
		file_antennas_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAntennaByIdResponse); i {
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
		file_antennas_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAntennaByIdRequest); i {
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
		file_antennas_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAntennaByIdResponse); i {
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
			RawDescriptor: file_antennas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_antennas_proto_goTypes,
		DependencyIndexes: file_antennas_proto_depIdxs,
		MessageInfos:      file_antennas_proto_msgTypes,
	}.Build()
	File_antennas_proto = out.File
	file_antennas_proto_rawDesc = nil
	file_antennas_proto_goTypes = nil
	file_antennas_proto_depIdxs = nil
}
