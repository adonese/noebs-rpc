// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package noebs_rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PaymentAPIClient is the client API for PaymentAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentAPIClient interface {
	// Sends a greeting
	GetWorkingKey(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetPurchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*Response, error)
	GetSpecialPayment(ctx context.Context, in *SpecialPaymentRequest, opts ...grpc.CallOption) (*Response, error)
	GetConsumerKey(ctx context.Context, in *ConsumerKeyRequest, opts ...grpc.CallOption) (*Response, error)
}

type paymentAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentAPIClient(cc grpc.ClientConnInterface) PaymentAPIClient {
	return &paymentAPIClient{cc}
}

func (c *paymentAPIClient) GetWorkingKey(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/noebs_rpc.PaymentAPI/GetWorkingKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentAPIClient) GetPurchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/noebs_rpc.PaymentAPI/GetPurchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentAPIClient) GetSpecialPayment(ctx context.Context, in *SpecialPaymentRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/noebs_rpc.PaymentAPI/GetSpecialPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentAPIClient) GetConsumerKey(ctx context.Context, in *ConsumerKeyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/noebs_rpc.PaymentAPI/GetConsumerKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentAPIServer is the server API for PaymentAPI service.
// All implementations must embed UnimplementedPaymentAPIServer
// for forward compatibility
type PaymentAPIServer interface {
	// Sends a greeting
	GetWorkingKey(context.Context, *Request) (*Response, error)
	GetPurchase(context.Context, *PurchaseRequest) (*Response, error)
	GetSpecialPayment(context.Context, *SpecialPaymentRequest) (*Response, error)
	GetConsumerKey(context.Context, *ConsumerKeyRequest) (*Response, error)
	mustEmbedUnimplementedPaymentAPIServer()
}

// UnimplementedPaymentAPIServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentAPIServer struct {
}

func (UnimplementedPaymentAPIServer) GetWorkingKey(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkingKey not implemented")
}
func (UnimplementedPaymentAPIServer) GetPurchase(context.Context, *PurchaseRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPurchase not implemented")
}
func (UnimplementedPaymentAPIServer) GetSpecialPayment(context.Context, *SpecialPaymentRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpecialPayment not implemented")
}
func (UnimplementedPaymentAPIServer) GetConsumerKey(context.Context, *ConsumerKeyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumerKey not implemented")
}
func (UnimplementedPaymentAPIServer) mustEmbedUnimplementedPaymentAPIServer() {}

// UnsafePaymentAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentAPIServer will
// result in compilation errors.
type UnsafePaymentAPIServer interface {
	mustEmbedUnimplementedPaymentAPIServer()
}

func RegisterPaymentAPIServer(s grpc.ServiceRegistrar, srv PaymentAPIServer) {
	s.RegisterService(&PaymentAPI_ServiceDesc, srv)
}

func _PaymentAPI_GetWorkingKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentAPIServer).GetWorkingKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noebs_rpc.PaymentAPI/GetWorkingKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentAPIServer).GetWorkingKey(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentAPI_GetPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentAPIServer).GetPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noebs_rpc.PaymentAPI/GetPurchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentAPIServer).GetPurchase(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentAPI_GetSpecialPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpecialPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentAPIServer).GetSpecialPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noebs_rpc.PaymentAPI/GetSpecialPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentAPIServer).GetSpecialPayment(ctx, req.(*SpecialPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentAPI_GetConsumerKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentAPIServer).GetConsumerKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noebs_rpc.PaymentAPI/GetConsumerKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentAPIServer).GetConsumerKey(ctx, req.(*ConsumerKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentAPI_ServiceDesc is the grpc.ServiceDesc for PaymentAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "noebs_rpc.PaymentAPI",
	HandlerType: (*PaymentAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkingKey",
			Handler:    _PaymentAPI_GetWorkingKey_Handler,
		},
		{
			MethodName: "GetPurchase",
			Handler:    _PaymentAPI_GetPurchase_Handler,
		},
		{
			MethodName: "GetSpecialPayment",
			Handler:    _PaymentAPI_GetSpecialPayment_Handler,
		},
		{
			MethodName: "GetConsumerKey",
			Handler:    _PaymentAPI_GetConsumerKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "key.proto",
}
