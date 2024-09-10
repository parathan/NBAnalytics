// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0--rc2
// source: predict.proto

package predictions

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PredictionService_Predict_FullMethodName = "/predictions.PredictionService/Predict"
)

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PredictionServiceClient interface {
	Predict(ctx context.Context, in *PredictionRequest, opts ...grpc.CallOption) (*PredictionResponse, error)
}

type predictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictionServiceClient(cc grpc.ClientConnInterface) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictionRequest, opts ...grpc.CallOption) (*PredictionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PredictionResponse)
	err := c.cc.Invoke(ctx, PredictionService_Predict_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
// All implementations must embed UnimplementedPredictionServiceServer
// for forward compatibility.
type PredictionServiceServer interface {
	Predict(context.Context, *PredictionRequest) (*PredictionResponse, error)
	mustEmbedUnimplementedPredictionServiceServer()
}

// UnimplementedPredictionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPredictionServiceServer struct{}

func (UnimplementedPredictionServiceServer) Predict(context.Context, *PredictionRequest) (*PredictionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (UnimplementedPredictionServiceServer) mustEmbedUnimplementedPredictionServiceServer() {}
func (UnimplementedPredictionServiceServer) testEmbeddedByValue()                           {}

// UnsafePredictionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PredictionServiceServer will
// result in compilation errors.
type UnsafePredictionServiceServer interface {
	mustEmbedUnimplementedPredictionServiceServer()
}

func RegisterPredictionServiceServer(s grpc.ServiceRegistrar, srv PredictionServiceServer) {
	// If the following call pancis, it indicates UnimplementedPredictionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PredictionService_ServiceDesc, srv)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PredictionService_Predict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PredictionService_ServiceDesc is the grpc.ServiceDesc for PredictionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PredictionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "predictions.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "predict.proto",
}
