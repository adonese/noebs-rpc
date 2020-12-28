// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: key.proto

package noebs

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TranDateTime string `protobuf:"bytes,1,opt,name=TranDateTime,proto3" json:"TranDateTime,omitempty"`
	TerminalID   string `protobuf:"bytes,2,opt,name=TerminalID,proto3" json:"TerminalID,omitempty"`
	ClientID     string `protobuf:"bytes,3,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	STAN         int32  `protobuf:"varint,4,opt,name=STAN,proto3" json:"STAN,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_key_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetTranDateTime() string {
	if x != nil {
		return x.TranDateTime
	}
	return ""
}

func (x *Request) GetTerminalID() string {
	if x != nil {
		return x.TerminalID
	}
	return ""
}

func (x *Request) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *Request) GetSTAN() int32 {
	if x != nil {
		return x.STAN
	}
	return 0
}

type PurchaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TranDateTime string  `protobuf:"bytes,1,opt,name=TranDateTime,proto3" json:"TranDateTime,omitempty"`
	TerminalID   string  `protobuf:"bytes,2,opt,name=TerminalID,proto3" json:"TerminalID,omitempty"`
	ClientID     string  `protobuf:"bytes,3,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	STAN         int32   `protobuf:"varint,4,opt,name=STAN,proto3" json:"STAN,omitempty"`
	Pan          string  `protobuf:"bytes,5,opt,name=Pan,proto3" json:"Pan,omitempty"`
	Expdate      string  `protobuf:"bytes,6,opt,name=Expdate,proto3" json:"Expdate,omitempty"`
	Pin          string  `protobuf:"bytes,7,opt,name=Pin,proto3" json:"Pin,omitempty"`
	Currency     string  `protobuf:"bytes,8,opt,name=Currency,proto3" json:"Currency,omitempty"`
	Amount       float32 `protobuf:"fixed32,9,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *PurchaseRequest) Reset() {
	*x = PurchaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseRequest) ProtoMessage() {}

func (x *PurchaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseRequest.ProtoReflect.Descriptor instead.
func (*PurchaseRequest) Descriptor() ([]byte, []int) {
	return file_key_proto_rawDescGZIP(), []int{1}
}

func (x *PurchaseRequest) GetTranDateTime() string {
	if x != nil {
		return x.TranDateTime
	}
	return ""
}

func (x *PurchaseRequest) GetTerminalID() string {
	if x != nil {
		return x.TerminalID
	}
	return ""
}

func (x *PurchaseRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *PurchaseRequest) GetSTAN() int32 {
	if x != nil {
		return x.STAN
	}
	return 0
}

func (x *PurchaseRequest) GetPan() string {
	if x != nil {
		return x.Pan
	}
	return ""
}

func (x *PurchaseRequest) GetExpdate() string {
	if x != nil {
		return x.Expdate
	}
	return ""
}

func (x *PurchaseRequest) GetPin() string {
	if x != nil {
		return x.Pin
	}
	return ""
}

func (x *PurchaseRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PurchaseRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SpecialPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TranDateTime  string  `protobuf:"bytes,1,opt,name=TranDateTime,proto3" json:"TranDateTime,omitempty"`
	UUID          string  `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	ApplicationID string  `protobuf:"bytes,3,opt,name=ApplicationID,proto3" json:"ApplicationID,omitempty"`
	Pan           string  `protobuf:"bytes,4,opt,name=Pan,proto3" json:"Pan,omitempty"`
	Expdate       string  `protobuf:"bytes,5,opt,name=Expdate,proto3" json:"Expdate,omitempty"`
	IPin          string  `protobuf:"bytes,6,opt,name=IPin,proto3" json:"IPin,omitempty"`
	Currency      string  `protobuf:"bytes,8,opt,name=Currency,proto3" json:"Currency,omitempty"`
	Amount        float32 `protobuf:"fixed32,9,opt,name=Amount,proto3" json:"Amount,omitempty"`
	BillerID      string  `protobuf:"bytes,10,opt,name=BillerID,proto3" json:"BillerID,omitempty"`
}

func (x *SpecialPaymentRequest) Reset() {
	*x = SpecialPaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecialPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecialPaymentRequest) ProtoMessage() {}

func (x *SpecialPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecialPaymentRequest.ProtoReflect.Descriptor instead.
func (*SpecialPaymentRequest) Descriptor() ([]byte, []int) {
	return file_key_proto_rawDescGZIP(), []int{2}
}

func (x *SpecialPaymentRequest) GetTranDateTime() string {
	if x != nil {
		return x.TranDateTime
	}
	return ""
}

func (x *SpecialPaymentRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *SpecialPaymentRequest) GetApplicationID() string {
	if x != nil {
		return x.ApplicationID
	}
	return ""
}

func (x *SpecialPaymentRequest) GetPan() string {
	if x != nil {
		return x.Pan
	}
	return ""
}

func (x *SpecialPaymentRequest) GetExpdate() string {
	if x != nil {
		return x.Expdate
	}
	return ""
}

func (x *SpecialPaymentRequest) GetIPin() string {
	if x != nil {
		return x.IPin
	}
	return ""
}

func (x *SpecialPaymentRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *SpecialPaymentRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *SpecialPaymentRequest) GetBillerID() string {
	if x != nil {
		return x.BillerID
	}
	return ""
}

// The response message containing the greetings
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkingKey      string `protobuf:"bytes,1,opt,name=WorkingKey,proto3" json:"WorkingKey,omitempty"`
	ResponseCode    int32  `protobuf:"varint,2,opt,name=ResponseCode,proto3" json:"ResponseCode,omitempty"`
	ResponseMessage string `protobuf:"bytes,3,opt,name=ResponseMessage,proto3" json:"ResponseMessage,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_key_proto_msgTypes[3]
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
	return file_key_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetWorkingKey() string {
	if x != nil {
		return x.WorkingKey
	}
	return ""
}

func (x *Response) GetResponseCode() int32 {
	if x != nil {
		return x.ResponseCode
	}
	return 0
}

func (x *Response) GetResponseMessage() string {
	if x != nil {
		return x.ResponseMessage
	}
	return ""
}

var File_key_proto protoreflect.FileDescriptor

var file_key_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x65,
	0x62, 0x73, 0x22, 0x7d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x54, 0x72, 0x61, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x54, 0x41, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x54, 0x41,
	0x4e, 0x22, 0xf7, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x72, 0x61,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x54, 0x41, 0x4e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x54, 0x41, 0x4e, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x61, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x50, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x45,
	0x78, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x78,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x50, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x85, 0x02, 0x0a, 0x15,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x72, 0x61,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x24, 0x0a,
	0x0d, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x50, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x78, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x49, 0x50, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49,
	0x50, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x69, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x44, 0x22, 0x78, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12,
	0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc0, 0x01,
	0x0a, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x50, 0x49, 0x12, 0x32, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x2e,
	0x6e, 0x6f, 0x65, 0x62, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x6e, 0x6f, 0x65, 0x62, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12,
	0x16, 0x2e, 0x6e, 0x6f, 0x65, 0x62, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6e, 0x6f, 0x65, 0x62, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x2e, 0x6e, 0x6f, 0x65, 0x62, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x6e, 0x6f, 0x65, 0x62, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x3f, 0x0a, 0x0c, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x42, 0x09, 0x52, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6f, 0x6e, 0x65, 0x73,
	0x65, 0x2f, 0x6e, 0x6f, 0x65, 0x62, 0x73, 0x2d, 0x72, 0x70, 0x63, 0x3b, 0x6e, 0x6f, 0x65, 0x62,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_key_proto_rawDescOnce sync.Once
	file_key_proto_rawDescData = file_key_proto_rawDesc
)

func file_key_proto_rawDescGZIP() []byte {
	file_key_proto_rawDescOnce.Do(func() {
		file_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_key_proto_rawDescData)
	})
	return file_key_proto_rawDescData
}

var file_key_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_key_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: noebs.Request
	(*PurchaseRequest)(nil),       // 1: noebs.PurchaseRequest
	(*SpecialPaymentRequest)(nil), // 2: noebs.SpecialPaymentRequest
	(*Response)(nil),              // 3: noebs.Response
}
var file_key_proto_depIdxs = []int32{
	0, // 0: noebs.PaymentAPI.GetWorkingKey:input_type -> noebs.Request
	1, // 1: noebs.PaymentAPI.GetPurchase:input_type -> noebs.PurchaseRequest
	2, // 2: noebs.PaymentAPI.GetSpecialPayment:input_type -> noebs.SpecialPaymentRequest
	3, // 3: noebs.PaymentAPI.GetWorkingKey:output_type -> noebs.Response
	3, // 4: noebs.PaymentAPI.GetPurchase:output_type -> noebs.Response
	3, // 5: noebs.PaymentAPI.GetSpecialPayment:output_type -> noebs.Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_key_proto_init() }
func file_key_proto_init() {
	if File_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseRequest); i {
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
		file_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecialPaymentRequest); i {
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
		file_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_key_proto_goTypes,
		DependencyIndexes: file_key_proto_depIdxs,
		MessageInfos:      file_key_proto_msgTypes,
	}.Build()
	File_key_proto = out.File
	file_key_proto_rawDesc = nil
	file_key_proto_goTypes = nil
	file_key_proto_depIdxs = nil
}
