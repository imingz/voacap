// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.26.1
// source: common.proto

package common

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

type Antenna struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AntennaID int64  `protobuf:"varint,1,opt,name=antennaID,proto3" json:"antennaID,omitempty" form:"antennaID" query:"antennaID"`
	Aname     string `protobuf:"bytes,2,opt,name=Aname,proto3" json:"Aname,omitempty" form:"Aname" query:"Aname"`
	Afile     string `protobuf:"bytes,3,opt,name=Afile,proto3" json:"Afile,omitempty" form:"Afile" query:"Afile"`
	AfbandMin int64  `protobuf:"varint,4,opt,name=AfbandMin,proto3" json:"AfbandMin,omitempty" form:"AfbandMin" query:"AfbandMin"`
	AfbandMax int64  `protobuf:"varint,5,opt,name=AfbandMax,proto3" json:"AfbandMax,omitempty" form:"AfbandMax" query:"AfbandMax"`
}

func (x *Antenna) Reset() {
	*x = Antenna{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Antenna) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Antenna) ProtoMessage() {}

func (x *Antenna) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Antenna.ProtoReflect.Descriptor instead.
func (*Antenna) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Antenna) GetAntennaID() int64 {
	if x != nil {
		return x.AntennaID
	}
	return 0
}

func (x *Antenna) GetAname() string {
	if x != nil {
		return x.Aname
	}
	return ""
}

func (x *Antenna) GetAfile() string {
	if x != nil {
		return x.Afile
	}
	return ""
}

func (x *Antenna) GetAfbandMin() int64 {
	if x != nil {
		return x.AfbandMin
	}
	return 0
}

func (x *Antenna) GetAfbandMax() int64 {
	if x != nil {
		return x.AfbandMax
	}
	return 0
}

type Station struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationID  int64   `protobuf:"varint,1,opt,name=stationID,proto3" json:"stationID,omitempty" form:"stationID" query:"stationID"`
	Slatitude  float64 `protobuf:"fixed64,2,opt,name=Slatitude,proto3" json:"Slatitude,omitempty" form:"Slatitude" query:"Slatitude"`
	Slongitude float64 `protobuf:"fixed64,3,opt,name=Slongitude,proto3" json:"Slongitude,omitempty" form:"Slongitude" query:"Slongitude"`
	Sname      string  `protobuf:"bytes,4,opt,name=Sname,proto3" json:"Sname,omitempty" form:"Sname" query:"Sname"`
}

func (x *Station) Reset() {
	*x = Station{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Station) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Station) ProtoMessage() {}

func (x *Station) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Station.ProtoReflect.Descriptor instead.
func (*Station) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Station) GetStationID() int64 {
	if x != nil {
		return x.StationID
	}
	return 0
}

func (x *Station) GetSlatitude() float64 {
	if x != nil {
		return x.Slatitude
	}
	return 0
}

func (x *Station) GetSlongitude() float64 {
	if x != nil {
		return x.Slongitude
	}
	return 0
}

func (x *Station) GetSname() string {
	if x != nil {
		return x.Sname
	}
	return ""
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CircuitReliability float64 `protobuf:"fixed64,1,opt,name=circuitReliability,proto3" json:"circuitReliability,omitempty" form:"circuitReliability" query:"circuitReliability"`
	Coefficient        string  `protobuf:"bytes,2,opt,name=coefficient,proto3" json:"coefficient,omitempty" form:"coefficient" query:"coefficient"`
	Date               string  `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty" form:"date" query:"date"`
	IxAntennaFbandMax  int64   `protobuf:"varint,4,opt,name=ixAntennaFbandMax,proto3" json:"ixAntennaFbandMax,omitempty" form:"ixAntennaFbandMax" query:"ixAntennaFbandMax"`
	IxAntennaFbandMin  int64   `protobuf:"varint,5,opt,name=ixAntennaFbandMin,proto3" json:"ixAntennaFbandMin,omitempty" form:"ixAntennaFbandMin" query:"ixAntennaFbandMin"`
	IxAntennaFile      string  `protobuf:"bytes,6,opt,name=ixAntennaFile,proto3" json:"ixAntennaFile,omitempty" form:"ixAntennaFile" query:"ixAntennaFile"`
	IxAntennaID        int64   `protobuf:"varint,7,opt,name=ixAntennaID,proto3" json:"ixAntennaID,omitempty" form:"ixAntennaID" query:"ixAntennaID"`
	IxAntennaName      string  `protobuf:"bytes,8,opt,name=ixAntennaName,proto3" json:"ixAntennaName,omitempty" form:"ixAntennaName" query:"ixAntennaName"`
	IxPower            float64 `protobuf:"fixed64,9,opt,name=ixPower,proto3" json:"ixPower,omitempty" form:"ixPower" query:"ixPower"`
	IxStationID        int64   `protobuf:"varint,10,opt,name=ixStationID,proto3" json:"ixStationID,omitempty" form:"ixStationID" query:"ixStationID"`
	IxStationLat       float64 `protobuf:"fixed64,11,opt,name=ixStationLat,proto3" json:"ixStationLat,omitempty" form:"ixStationLat" query:"ixStationLat"`
	IxStationLng       float64 `protobuf:"fixed64,12,opt,name=ixStationLng,proto3" json:"ixStationLng,omitempty" form:"ixStationLng" query:"ixStationLng"`
	IxStationName      string  `protobuf:"bytes,13,opt,name=ixStationName,proto3" json:"ixStationName,omitempty" form:"ixStationName" query:"ixStationName"`
	LinkID             int64   `protobuf:"varint,14,opt,name=linkID,proto3" json:"linkID,omitempty" form:"linkID" query:"linkID"`
	LinkType           string  `protobuf:"bytes,15,opt,name=linkType,proto3" json:"linkType,omitempty" form:"linkType" query:"linkType"`
	Noise              float64 `protobuf:"fixed64,16,opt,name=noise,proto3" json:"noise,omitempty" form:"noise" query:"noise"`
	RxAntennaFbandMax  int64   `protobuf:"varint,17,opt,name=rxAntennaFbandMax,proto3" json:"rxAntennaFbandMax,omitempty" form:"rxAntennaFbandMax" query:"rxAntennaFbandMax"`
	RxAntennaFbandMin  int64   `protobuf:"varint,18,opt,name=rxAntennaFbandMin,proto3" json:"rxAntennaFbandMin,omitempty" form:"rxAntennaFbandMin" query:"rxAntennaFbandMin"`
	RxAntennaFile      string  `protobuf:"bytes,19,opt,name=rxAntennaFile,proto3" json:"rxAntennaFile,omitempty" form:"rxAntennaFile" query:"rxAntennaFile"`
	RxAntennaID        int64   `protobuf:"varint,20,opt,name=rxAntennaID,proto3" json:"rxAntennaID,omitempty" form:"rxAntennaID" query:"rxAntennaID"`
	RxAntennaName      string  `protobuf:"bytes,21,opt,name=rxAntennaName,proto3" json:"rxAntennaName,omitempty" form:"rxAntennaName" query:"rxAntennaName"`
	RxStationID        int64   `protobuf:"varint,22,opt,name=rxStationID,proto3" json:"rxStationID,omitempty" form:"rxStationID" query:"rxStationID"`
	RxStationLat       float64 `protobuf:"fixed64,23,opt,name=rxStationLat,proto3" json:"rxStationLat,omitempty" form:"rxStationLat" query:"rxStationLat"`
	RxStationLng       float64 `protobuf:"fixed64,24,opt,name=rxStationLng,proto3" json:"rxStationLng,omitempty" form:"rxStationLng" query:"rxStationLng"`
	RxStationName      string  `protobuf:"bytes,25,opt,name=rxStationName,proto3" json:"rxStationName,omitempty" form:"rxStationName" query:"rxStationName"`
	SNR                float64 `protobuf:"fixed64,26,opt,name=SNR,proto3" json:"SNR,omitempty" form:"SNR" query:"SNR"`
	SunspotNum         int64   `protobuf:"varint,27,opt,name=sunspotNum,proto3" json:"sunspotNum,omitempty" form:"sunspotNum" query:"sunspotNum"`
	TimeType           string  `protobuf:"bytes,28,opt,name=timeType,proto3" json:"timeType,omitempty" form:"timeType" query:"timeType"`
	TxAntennaFbandMax  int64   `protobuf:"varint,29,opt,name=txAntennaFbandMax,proto3" json:"txAntennaFbandMax,omitempty" form:"txAntennaFbandMax" query:"txAntennaFbandMax"`
	TxAntennaFbandMin  int64   `protobuf:"varint,30,opt,name=txAntennaFbandMin,proto3" json:"txAntennaFbandMin,omitempty" form:"txAntennaFbandMin" query:"txAntennaFbandMin"`
	TxAntennaFile      string  `protobuf:"bytes,31,opt,name=txAntennaFile,proto3" json:"txAntennaFile,omitempty" form:"txAntennaFile" query:"txAntennaFile"`
	TxAntennaID        int64   `protobuf:"varint,32,opt,name=txAntennaID,proto3" json:"txAntennaID,omitempty" form:"txAntennaID" query:"txAntennaID"`
	TxAntennaName      string  `protobuf:"bytes,33,opt,name=txAntennaName,proto3" json:"txAntennaName,omitempty" form:"txAntennaName" query:"txAntennaName"`
	TxPower            float64 `protobuf:"fixed64,34,opt,name=txPower,proto3" json:"txPower,omitempty" form:"txPower" query:"txPower"`
	TxStationID        int64   `protobuf:"varint,35,opt,name=txStationID,proto3" json:"txStationID,omitempty" form:"txStationID" query:"txStationID"`
	TxStationLat       float64 `protobuf:"fixed64,36,opt,name=txStationLat,proto3" json:"txStationLat,omitempty" form:"txStationLat" query:"txStationLat"`
	TxStationLng       float64 `protobuf:"fixed64,37,opt,name=txStationLng,proto3" json:"txStationLng,omitempty" form:"txStationLng" query:"txStationLng"`
	TxStationName      string  `protobuf:"bytes,38,opt,name=txStationName,proto3" json:"txStationName,omitempty" form:"txStationName" query:"txStationName"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Link) GetCircuitReliability() float64 {
	if x != nil {
		return x.CircuitReliability
	}
	return 0
}

func (x *Link) GetCoefficient() string {
	if x != nil {
		return x.Coefficient
	}
	return ""
}

func (x *Link) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Link) GetIxAntennaFbandMax() int64 {
	if x != nil {
		return x.IxAntennaFbandMax
	}
	return 0
}

func (x *Link) GetIxAntennaFbandMin() int64 {
	if x != nil {
		return x.IxAntennaFbandMin
	}
	return 0
}

func (x *Link) GetIxAntennaFile() string {
	if x != nil {
		return x.IxAntennaFile
	}
	return ""
}

func (x *Link) GetIxAntennaID() int64 {
	if x != nil {
		return x.IxAntennaID
	}
	return 0
}

func (x *Link) GetIxAntennaName() string {
	if x != nil {
		return x.IxAntennaName
	}
	return ""
}

func (x *Link) GetIxPower() float64 {
	if x != nil {
		return x.IxPower
	}
	return 0
}

func (x *Link) GetIxStationID() int64 {
	if x != nil {
		return x.IxStationID
	}
	return 0
}

func (x *Link) GetIxStationLat() float64 {
	if x != nil {
		return x.IxStationLat
	}
	return 0
}

func (x *Link) GetIxStationLng() float64 {
	if x != nil {
		return x.IxStationLng
	}
	return 0
}

func (x *Link) GetIxStationName() string {
	if x != nil {
		return x.IxStationName
	}
	return ""
}

func (x *Link) GetLinkID() int64 {
	if x != nil {
		return x.LinkID
	}
	return 0
}

func (x *Link) GetLinkType() string {
	if x != nil {
		return x.LinkType
	}
	return ""
}

func (x *Link) GetNoise() float64 {
	if x != nil {
		return x.Noise
	}
	return 0
}

func (x *Link) GetRxAntennaFbandMax() int64 {
	if x != nil {
		return x.RxAntennaFbandMax
	}
	return 0
}

func (x *Link) GetRxAntennaFbandMin() int64 {
	if x != nil {
		return x.RxAntennaFbandMin
	}
	return 0
}

func (x *Link) GetRxAntennaFile() string {
	if x != nil {
		return x.RxAntennaFile
	}
	return ""
}

func (x *Link) GetRxAntennaID() int64 {
	if x != nil {
		return x.RxAntennaID
	}
	return 0
}

func (x *Link) GetRxAntennaName() string {
	if x != nil {
		return x.RxAntennaName
	}
	return ""
}

func (x *Link) GetRxStationID() int64 {
	if x != nil {
		return x.RxStationID
	}
	return 0
}

func (x *Link) GetRxStationLat() float64 {
	if x != nil {
		return x.RxStationLat
	}
	return 0
}

func (x *Link) GetRxStationLng() float64 {
	if x != nil {
		return x.RxStationLng
	}
	return 0
}

func (x *Link) GetRxStationName() string {
	if x != nil {
		return x.RxStationName
	}
	return ""
}

func (x *Link) GetSNR() float64 {
	if x != nil {
		return x.SNR
	}
	return 0
}

func (x *Link) GetSunspotNum() int64 {
	if x != nil {
		return x.SunspotNum
	}
	return 0
}

func (x *Link) GetTimeType() string {
	if x != nil {
		return x.TimeType
	}
	return ""
}

func (x *Link) GetTxAntennaFbandMax() int64 {
	if x != nil {
		return x.TxAntennaFbandMax
	}
	return 0
}

func (x *Link) GetTxAntennaFbandMin() int64 {
	if x != nil {
		return x.TxAntennaFbandMin
	}
	return 0
}

func (x *Link) GetTxAntennaFile() string {
	if x != nil {
		return x.TxAntennaFile
	}
	return ""
}

func (x *Link) GetTxAntennaID() int64 {
	if x != nil {
		return x.TxAntennaID
	}
	return 0
}

func (x *Link) GetTxAntennaName() string {
	if x != nil {
		return x.TxAntennaName
	}
	return ""
}

func (x *Link) GetTxPower() float64 {
	if x != nil {
		return x.TxPower
	}
	return 0
}

func (x *Link) GetTxStationID() int64 {
	if x != nil {
		return x.TxStationID
	}
	return 0
}

func (x *Link) GetTxStationLat() float64 {
	if x != nil {
		return x.TxStationLat
	}
	return 0
}

func (x *Link) GetTxStationLng() float64 {
	if x != nil {
		return x.TxStationLng
	}
	return 0
}

func (x *Link) GetTxStationName() string {
	if x != nil {
		return x.TxStationName
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x8f, 0x01, 0x0a, 0x07, 0x41, 0x6e, 0x74, 0x65, 0x6e,
	0x6e, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x41, 0x66, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x41, 0x66, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x66,
	0x62, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41,
	0x66, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x22, 0x7b, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x53, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x53, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x53, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x53, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc6, 0x0a, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2e,
	0x0a, 0x12, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x52, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x63, 0x69, 0x72, 0x63,
	0x75, 0x69, 0x74, 0x52, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e,
	0x61, 0x46, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x69, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x62, 0x61, 0x6e, 0x64, 0x4d,
	0x61, 0x78, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46,
	0x62, 0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x69,
	0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e,
	0x12, 0x24, 0x0a, 0x0d, 0x69, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x69, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e,
	0x6e, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x78, 0x41, 0x6e, 0x74, 0x65,
	0x6e, 0x6e, 0x61, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x78, 0x41,
	0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x78, 0x41, 0x6e,
	0x74, 0x65, 0x6e, 0x6e, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x69, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x78, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x07, 0x69, 0x78, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x78, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69,
	0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x78,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0c, 0x69, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x69, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6e, 0x67, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x69, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x78, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x6b,
	0x49, 0x44, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x6b, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x69, 0x73, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6e, 0x6f, 0x69,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46,
	0x62, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x72,
	0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78,
	0x12, 0x2c, 0x0a, 0x11, 0x72, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x62, 0x61,
	0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x72, 0x78, 0x41,
	0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x12, 0x24,
	0x0a, 0x0d, 0x72, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e,
	0x61, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x78, 0x41, 0x6e, 0x74,
	0x65, 0x6e, 0x6e, 0x61, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x78, 0x41, 0x6e, 0x74, 0x65,
	0x6e, 0x6e, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72,
	0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x72, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x72, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x74, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x72, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6e, 0x67, 0x18, 0x18, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x72, 0x78, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x78, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72,
	0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x53, 0x4e, 0x52, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x53, 0x4e, 0x52, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x75, 0x6e, 0x73, 0x70, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x73, 0x75, 0x6e, 0x73, 0x70, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x78,
	0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x74, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61,
	0x46, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x78, 0x41, 0x6e,
	0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x62, 0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x74, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x62,
	0x61, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x78, 0x41, 0x6e, 0x74, 0x65,
	0x6e, 0x6e, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74,
	0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x49, 0x44, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x74, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x49, 0x44, 0x12, 0x24,
	0x0a, 0x0d, 0x74, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x78, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x78, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x18,
	0x22, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x74, 0x78, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x23, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x74,
	0x18, 0x24, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6e, 0x67, 0x18, 0x25, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x78, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x78, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x78, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x19,
	0x5a, 0x17, 0x76, 0x6f, 0x61, 0x63, 0x61, 0x70, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_proto_goTypes = []interface{}{
	(*Antenna)(nil), // 0: common.Antenna
	(*Station)(nil), // 1: common.Station
	(*Link)(nil),    // 2: common.Link
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Antenna); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Station); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
