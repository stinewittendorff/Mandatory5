// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.0--rc1
// source: proto/proto.proto

package proto

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

type BidResponse_RESULT int32

const (
	BidResponse_SUCCESS   BidResponse_RESULT = 0
	BidResponse_FAIL      BidResponse_RESULT = 1
	BidResponse_EXCEPTION BidResponse_RESULT = 2
)

// Enum value maps for BidResponse_RESULT.
var (
	BidResponse_RESULT_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
		2: "EXCEPTION",
	}
	BidResponse_RESULT_value = map[string]int32{
		"SUCCESS":   0,
		"FAIL":      1,
		"EXCEPTION": 2,
	}
)

func (x BidResponse_RESULT) Enum() *BidResponse_RESULT {
	p := new(BidResponse_RESULT)
	*p = x
	return p
}

func (x BidResponse_RESULT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BidResponse_RESULT) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_proto_proto_enumTypes[0].Descriptor()
}

func (BidResponse_RESULT) Type() protoreflect.EnumType {
	return &file_proto_proto_proto_enumTypes[0]
}

func (x BidResponse_RESULT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BidResponse_RESULT.Descriptor instead.
func (BidResponse_RESULT) EnumDescriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{4, 0}
}

type ResultResponse_STATE int32

const (
	ResultResponse_ONGOING    ResultResponse_STATE = 0
	ResultResponse_FINISHED   ResultResponse_STATE = 1
	ResultResponse_NOTSTARTED ResultResponse_STATE = 2
)

// Enum value maps for ResultResponse_STATE.
var (
	ResultResponse_STATE_name = map[int32]string{
		0: "ONGOING",
		1: "FINISHED",
		2: "NOTSTARTED",
	}
	ResultResponse_STATE_value = map[string]int32{
		"ONGOING":    0,
		"FINISHED":   1,
		"NOTSTARTED": 2,
	}
)

func (x ResultResponse_STATE) Enum() *ResultResponse_STATE {
	p := new(ResultResponse_STATE)
	*p = x
	return p
}

func (x ResultResponse_STATE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultResponse_STATE) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_proto_proto_enumTypes[1].Descriptor()
}

func (ResultResponse_STATE) Type() protoreflect.EnumType {
	return &file_proto_proto_proto_enumTypes[1]
}

func (x ResultResponse_STATE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultResponse_STATE.Descriptor instead.
func (ResultResponse_STATE) EnumDescriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{6, 0}
}

type ServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsPrimary bool `protobuf:"varint,1,opt,name=isPrimary,proto3" json:"isPrimary,omitempty"`
}

func (x *ServerResponse) Reset() {
	*x = ServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerResponse) ProtoMessage() {}

func (x *ServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerResponse.ProtoReflect.Descriptor instead.
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{0}
}

func (x *ServerResponse) GetIsPrimary() bool {
	if x != nil {
		return x.IsPrimary
	}
	return false
}

type PingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PingMessage) Reset() {
	*x = PingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingMessage) ProtoMessage() {}

func (x *PingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingMessage.ProtoReflect.Descriptor instead.
func (*PingMessage) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{1}
}

func (x *PingMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{2}
}

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lamport  int32 `protobuf:"varint,1,opt,name=Lamport,proto3" json:"Lamport,omitempty"`
	BidderId int32 `protobuf:"varint,2,opt,name=bidderId,proto3" json:"bidderId,omitempty"`
	Amount   int32 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{3}
}

func (x *BidRequest) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

func (x *BidRequest) GetBidderId() int32 {
	if x != nil {
		return x.BidderId
	}
	return 0
}

func (x *BidRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lamport int32              `protobuf:"varint,1,opt,name=Lamport,proto3" json:"Lamport,omitempty"`
	Result  BidResponse_RESULT `protobuf:"varint,2,opt,name=result,proto3,enum=BidResponse_RESULT" json:"result,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{4}
}

func (x *BidResponse) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

func (x *BidResponse) GetResult() BidResponse_RESULT {
	if x != nil {
		return x.Result
	}
	return BidResponse_SUCCESS
}

type ResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResultRequest) Reset() {
	*x = ResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRequest) ProtoMessage() {}

func (x *ResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRequest.ProtoReflect.Descriptor instead.
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{5}
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outcome       ResultResponse_STATE `protobuf:"varint,1,opt,name=outcome,proto3,enum=ResultResponse_STATE" json:"outcome,omitempty"` // "Auction ongoing" or winner's name
	HighestBid    int32                `protobuf:"varint,2,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	HighestBidder int32                `protobuf:"varint,3,opt,name=highestBidder,proto3" json:"highestBidder,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{6}
}

func (x *ResultResponse) GetOutcome() ResultResponse_STATE {
	if x != nil {
		return x.Outcome
	}
	return ResultResponse_ONGOING
}

func (x *ResultResponse) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *ResultResponse) GetHighestBidder() int32 {
	if x != nil {
		return x.HighestBidder
	}
	return 0
}

var File_proto_proto_proto protoreflect.FileDescriptor

var file_proto_proto_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x22, 0x1d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5a, 0x0a, 0x0a, 0x42,
	0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4c, 0x61, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2e,
	0x0a, 0x06, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0x0f,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xbb, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x54, 0x41, 0x54, 0x45, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x63,
	0x6f, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74,
	0x42, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x68, 0x69, 0x67, 0x68,
	0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x05, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x4e, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x4f, 0x54, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0x56, 0x0a,
	0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12,
	0x0b, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x42,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xaf, 0x01, 0x0a, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x42, 0x69, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x12, 0x0b, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x0e, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0c, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x0c, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x24, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x48, 0x61, 0x6e, 0x64, 0x69,
	0x6e, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proto_proto_rawDescOnce sync.Once
	file_proto_proto_proto_rawDescData = file_proto_proto_proto_rawDesc
)

func file_proto_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_proto_rawDescData)
	})
	return file_proto_proto_proto_rawDescData
}

var file_proto_proto_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_proto_proto_goTypes = []any{
	(BidResponse_RESULT)(0),   // 0: BidResponse.RESULT
	(ResultResponse_STATE)(0), // 1: ResultResponse.STATE
	(*ServerResponse)(nil),    // 2: ServerResponse
	(*PingMessage)(nil),       // 3: PingMessage
	(*Empty)(nil),             // 4: Empty
	(*BidRequest)(nil),        // 5: BidRequest
	(*BidResponse)(nil),       // 6: BidResponse
	(*ResultRequest)(nil),     // 7: ResultRequest
	(*ResultResponse)(nil),    // 8: ResultResponse
}
var file_proto_proto_proto_depIdxs = []int32{
	0, // 0: BidResponse.result:type_name -> BidResponse.RESULT
	1, // 1: ResultResponse.outcome:type_name -> ResultResponse.STATE
	5, // 2: Auction.Bid:input_type -> BidRequest
	7, // 3: Auction.Result:input_type -> ResultRequest
	5, // 4: backupNode.BidBackup:input_type -> BidRequest
	7, // 5: backupNode.ResultBackup:input_type -> ResultRequest
	3, // 6: backupNode.Ping:input_type -> PingMessage
	4, // 7: backupNode.isPrimary:input_type -> Empty
	6, // 8: Auction.Bid:output_type -> BidResponse
	8, // 9: Auction.Result:output_type -> ResultResponse
	6, // 10: backupNode.BidBackup:output_type -> BidResponse
	8, // 11: backupNode.ResultBackup:output_type -> ResultResponse
	3, // 12: backupNode.Ping:output_type -> PingMessage
	2, // 13: backupNode.isPrimary:output_type -> ServerResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_proto_proto_init() }
func file_proto_proto_proto_init() {
	if File_proto_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ServerResponse); i {
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
		file_proto_proto_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PingMessage); i {
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
		file_proto_proto_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_proto_proto_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BidRequest); i {
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
		file_proto_proto_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BidResponse); i {
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
		file_proto_proto_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ResultRequest); i {
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
		file_proto_proto_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ResultResponse); i {
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
			RawDescriptor: file_proto_proto_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_proto_depIdxs,
		EnumInfos:         file_proto_proto_proto_enumTypes,
		MessageInfos:      file_proto_proto_proto_msgTypes,
	}.Build()
	File_proto_proto_proto = out.File
	file_proto_proto_proto_rawDesc = nil
	file_proto_proto_proto_goTypes = nil
	file_proto_proto_proto_depIdxs = nil
}