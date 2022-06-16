//
// This protobuf file contains the definitions of all the calls and messages
// used by drand nodes themselves to create the distributed randomness beacon.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: crypto/dkg/dkg.proto

package dkg

import (
	common "github.com/drand/drand/protobuf/common"
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

// Packet is a wrapper around the three different types of DKG messages
type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Bundle:
	//	*Packet_Deal
	//	*Packet_Response
	//	*Packet_Justification
	Bundle   isPacket_Bundle  `protobuf_oneof:"Bundle"`
	Metadata *common.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_dkg_dkg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_dkg_dkg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_crypto_dkg_dkg_proto_rawDescGZIP(), []int{0}
}

func (m *Packet) GetBundle() isPacket_Bundle {
	if m != nil {
		return m.Bundle
	}
	return nil
}

func (x *Packet) GetDeal() *DealBundle {
	if x, ok := x.GetBundle().(*Packet_Deal); ok {
		return x.Deal
	}
	return nil
}

func (x *Packet) GetResponse() *ResponseBundle {
	if x, ok := x.GetBundle().(*Packet_Response); ok {
		return x.Response
	}
	return nil
}

func (x *Packet) GetJustification() *JustificationBundle {
	if x, ok := x.GetBundle().(*Packet_Justification); ok {
		return x.Justification
	}
	return nil
}

func (x *Packet) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type isPacket_Bundle interface {
	isPacket_Bundle()
}

type Packet_Deal struct {
	Deal *DealBundle `protobuf:"bytes,1,opt,name=deal,proto3,oneof"`
}

type Packet_Response struct {
	Response *ResponseBundle `protobuf:"bytes,2,opt,name=response,proto3,oneof"`
}

type Packet_Justification struct {
	Justification *JustificationBundle `protobuf:"bytes,3,opt,name=justification,proto3,oneof"`
}

func (*Packet_Deal) isPacket_Bundle() {}

func (*Packet_Response) isPacket_Bundle() {}

func (*Packet_Justification) isPacket_Bundle() {}

// DealBundle is a packet issued by a dealer that contains each individual
// deals, as well as the coefficients of the public polynomial he used.
type DealBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index of the dealer that issues these deals
	DealerIndex uint32 `protobuf:"varint,1,opt,name=dealer_index,json=dealerIndex,proto3" json:"dealer_index,omitempty"`
	// Coefficients of the public polynomial that is created from the
	// private polynomial from which the shares are derived.
	Commits [][]byte `protobuf:"bytes,2,rep,name=commits,proto3" json:"commits,omitempty"`
	// list of deals for each individual share holders.
	Deals []*Deal `protobuf:"bytes,3,rep,name=deals,proto3" json:"deals,omitempty"`
	// session identifier of the protocol run
	SessionId []byte `protobuf:"bytes,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// signature over the hash of the deal
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *DealBundle) Reset() {
	*x = DealBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_dkg_dkg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DealBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DealBundle) ProtoMessage() {}

func (x *DealBundle) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_dkg_dkg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DealBundle.ProtoReflect.Descriptor instead.
func (*DealBundle) Descriptor() ([]byte, []int) {
	return file_crypto_dkg_dkg_proto_rawDescGZIP(), []int{1}
}

func (x *DealBundle) GetDealerIndex() uint32 {
	if x != nil {
		return x.DealerIndex
	}
	return 0
}

func (x *DealBundle) GetCommits() [][]byte {
	if x != nil {
		return x.Commits
	}
	return nil
}

func (x *DealBundle) GetDeals() []*Deal {
	if x != nil {
		return x.Deals
	}
	return nil
}

func (x *DealBundle) GetSessionId() []byte {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *DealBundle) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// Deal contains a share for a participant.
type Deal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShareIndex uint32 `protobuf:"varint,1,opt,name=share_index,json=shareIndex,proto3" json:"share_index,omitempty"`
	// encryption of the share using ECIES
	EncryptedShare []byte `protobuf:"bytes,2,opt,name=encrypted_share,json=encryptedShare,proto3" json:"encrypted_share,omitempty"`
}

func (x *Deal) Reset() {
	*x = Deal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_dkg_dkg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deal) ProtoMessage() {}

func (x *Deal) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_dkg_dkg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deal.ProtoReflect.Descriptor instead.
func (*Deal) Descriptor() ([]byte, []int) {
	return file_crypto_dkg_dkg_proto_rawDescGZIP(), []int{2}
}

func (x *Deal) GetShareIndex() uint32 {
	if x != nil {
		return x.ShareIndex
	}
	return 0
}

func (x *Deal) GetEncryptedShare() []byte {
	if x != nil {
		return x.EncryptedShare
	}
	return nil
}

// ResponseBundle is a packet issued by a share holder that contains all the
// responses (complaint and/or success) to broadcast.
type ResponseBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShareIndex uint32      `protobuf:"varint,1,opt,name=share_index,json=shareIndex,proto3" json:"share_index,omitempty"`
	Responses  []*Response `protobuf:"bytes,2,rep,name=responses,proto3" json:"responses,omitempty"`
	// session identifier of the protocol run
	SessionId []byte `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// signature over the hash of the response
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ResponseBundle) Reset() {
	*x = ResponseBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_dkg_dkg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBundle) ProtoMessage() {}

func (x *ResponseBundle) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_dkg_dkg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBundle.ProtoReflect.Descriptor instead.
func (*ResponseBundle) Descriptor() ([]byte, []int) {
	return file_crypto_dkg_dkg_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseBundle) GetShareIndex() uint32 {
	if x != nil {
		return x.ShareIndex
	}
	return 0
}

func (x *ResponseBundle) GetResponses() []*Response {
	if x != nil {
		return x.Responses
	}
	return nil
}

func (x *ResponseBundle) GetSessionId() []byte {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *ResponseBundle) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// Response holds the response that a participant broadcast after having
// received a deal.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// index of the dealer for which this response is for
	DealerIndex uint32 `protobuf:"varint,1,opt,name=dealer_index,json=dealerIndex,proto3" json:"dealer_index,omitempty"`
	// Status represents a complaint if set to false, a success if set to
	// true.
	Status bool `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_dkg_dkg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_dkg_dkg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_crypto_dkg_dkg_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetDealerIndex() uint32 {
	if x != nil {
		return x.DealerIndex
	}
	return 0
}

func (x *Response) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

// JustificationBundle is a packet that holds all justifications a dealer must
// produce
type JustificationBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DealerIndex    uint32           `protobuf:"varint,1,opt,name=dealer_index,json=dealerIndex,proto3" json:"dealer_index,omitempty"`
	Justifications []*Justification `protobuf:"bytes,2,rep,name=justifications,proto3" json:"justifications,omitempty"`
	// session identifier of the protocol run
	SessionId []byte `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// signature over the hash of the justification
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *JustificationBundle) Reset() {
	*x = JustificationBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_dkg_dkg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JustificationBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JustificationBundle) ProtoMessage() {}

func (x *JustificationBundle) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_dkg_dkg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JustificationBundle.ProtoReflect.Descriptor instead.
func (*JustificationBundle) Descriptor() ([]byte, []int) {
	return file_crypto_dkg_dkg_proto_rawDescGZIP(), []int{5}
}

func (x *JustificationBundle) GetDealerIndex() uint32 {
	if x != nil {
		return x.DealerIndex
	}
	return 0
}

func (x *JustificationBundle) GetJustifications() []*Justification {
	if x != nil {
		return x.Justifications
	}
	return nil
}

func (x *JustificationBundle) GetSessionId() []byte {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *JustificationBundle) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// Justification holds the justification from a dealer after a participant
// issued a complaint response because of a supposedly invalid deal.
type Justification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// represents for who share holder this justification is
	ShareIndex uint32 `protobuf:"varint,1,opt,name=share_index,json=shareIndex,proto3" json:"share_index,omitempty"`
	// plaintext share so everyone can see it correct
	Share []byte `protobuf:"bytes,2,opt,name=share,proto3" json:"share,omitempty"`
}

func (x *Justification) Reset() {
	*x = Justification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_dkg_dkg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Justification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Justification) ProtoMessage() {}

func (x *Justification) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_dkg_dkg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Justification.ProtoReflect.Descriptor instead.
func (*Justification) Descriptor() ([]byte, []int) {
	return file_crypto_dkg_dkg_proto_rawDescGZIP(), []int{6}
}

func (x *Justification) GetShareIndex() uint32 {
	if x != nil {
		return x.ShareIndex
	}
	return 0
}

func (x *Justification) GetShare() []byte {
	if x != nil {
		return x.Share
	}
	return nil
}

var File_crypto_dkg_dkg_proto protoreflect.FileDescriptor

var file_crypto_dkg_dkg_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x64, 0x6b, 0x67, 0x2f, 0x64, 0x6b, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x64, 0x6b, 0x67, 0x1a, 0x13, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdc, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x64,
	0x65, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x6b, 0x67, 0x2e,
	0x44, 0x65, 0x61, 0x6c, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x04, 0x64, 0x65,
	0x61, 0x6c, 0x12, 0x31, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x6b, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64,
	0x6b, 0x67, 0x2e, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x08, 0x0a, 0x06, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22,
	0xa7, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x6c, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x64,
	0x65, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64, 0x6b, 0x67,
	0x2e, 0x44, 0x65, 0x61, 0x6c, 0x52, 0x05, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x50, 0x0a, 0x04, 0x44, 0x65, 0x61,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x2b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x6b, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x45, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x65, 0x61,
	0x6c, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xb1, 0x01, 0x0a, 0x13, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x61, 0x6c,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x3a, 0x0a, 0x0e, 0x6a,
	0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6b, 0x67, 0x2e, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x46, 0x0a, 0x0d, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x61, 0x6e, 0x64,
	0x2f, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x64, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_crypto_dkg_dkg_proto_rawDescOnce sync.Once
	file_crypto_dkg_dkg_proto_rawDescData = file_crypto_dkg_dkg_proto_rawDesc
)

func file_crypto_dkg_dkg_proto_rawDescGZIP() []byte {
	file_crypto_dkg_dkg_proto_rawDescOnce.Do(func() {
		file_crypto_dkg_dkg_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_dkg_dkg_proto_rawDescData)
	})
	return file_crypto_dkg_dkg_proto_rawDescData
}

var file_crypto_dkg_dkg_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_crypto_dkg_dkg_proto_goTypes = []interface{}{
	(*Packet)(nil),              // 0: dkg.Packet
	(*DealBundle)(nil),          // 1: dkg.DealBundle
	(*Deal)(nil),                // 2: dkg.Deal
	(*ResponseBundle)(nil),      // 3: dkg.ResponseBundle
	(*Response)(nil),            // 4: dkg.Response
	(*JustificationBundle)(nil), // 5: dkg.JustificationBundle
	(*Justification)(nil),       // 6: dkg.Justification
	(*common.Metadata)(nil),     // 7: common.Metadata
}
var file_crypto_dkg_dkg_proto_depIdxs = []int32{
	1, // 0: dkg.Packet.deal:type_name -> dkg.DealBundle
	3, // 1: dkg.Packet.response:type_name -> dkg.ResponseBundle
	5, // 2: dkg.Packet.justification:type_name -> dkg.JustificationBundle
	7, // 3: dkg.Packet.metadata:type_name -> common.Metadata
	2, // 4: dkg.DealBundle.deals:type_name -> dkg.Deal
	4, // 5: dkg.ResponseBundle.responses:type_name -> dkg.Response
	6, // 6: dkg.JustificationBundle.justifications:type_name -> dkg.Justification
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_crypto_dkg_dkg_proto_init() }
func file_crypto_dkg_dkg_proto_init() {
	if File_crypto_dkg_dkg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crypto_dkg_dkg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
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
		file_crypto_dkg_dkg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DealBundle); i {
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
		file_crypto_dkg_dkg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deal); i {
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
		file_crypto_dkg_dkg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBundle); i {
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
		file_crypto_dkg_dkg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_crypto_dkg_dkg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JustificationBundle); i {
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
		file_crypto_dkg_dkg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Justification); i {
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
	file_crypto_dkg_dkg_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Packet_Deal)(nil),
		(*Packet_Response)(nil),
		(*Packet_Justification)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crypto_dkg_dkg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crypto_dkg_dkg_proto_goTypes,
		DependencyIndexes: file_crypto_dkg_dkg_proto_depIdxs,
		MessageInfos:      file_crypto_dkg_dkg_proto_msgTypes,
	}.Build()
	File_crypto_dkg_dkg_proto = out.File
	file_crypto_dkg_dkg_proto_rawDesc = nil
	file_crypto_dkg_dkg_proto_goTypes = nil
	file_crypto_dkg_dkg_proto_depIdxs = nil
}
