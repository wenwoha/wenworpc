// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: wenworpc/scrapy.proto

package scrapy

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ScrapyClient is the client API for Scrapy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScrapyClient interface {
	GetArticles(ctx context.Context, in *ArticleScrapyRequest, opts ...grpc.CallOption) (*ArticleScrapyResponse, error)
}

type scrapyClient struct {
	cc grpc.ClientConnInterface
}

func NewScrapyClient(cc grpc.ClientConnInterface) ScrapyClient {
	return &scrapyClient{cc}
}

func (c *scrapyClient) GetArticles(ctx context.Context, in *ArticleScrapyRequest, opts ...grpc.CallOption) (*ArticleScrapyResponse, error) {
	out := new(ArticleScrapyResponse)
	err := c.cc.Invoke(ctx, "/wenworpc.scrapy.Scrapy/getArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScrapyServer is the server API for Scrapy service.
// All implementations must embed UnimplementedScrapyServer
// for forward compatibility
type ScrapyServer interface {
	GetArticles(context.Context, *ArticleScrapyRequest) (*ArticleScrapyResponse, error)
	mustEmbedUnimplementedScrapyServer()
}

// UnimplementedScrapyServer must be embedded to have forward compatible implementations.
type UnimplementedScrapyServer struct {
}

func (UnimplementedScrapyServer) GetArticles(context.Context, *ArticleScrapyRequest) (*ArticleScrapyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedScrapyServer) mustEmbedUnimplementedScrapyServer() {}

// UnsafeScrapyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScrapyServer will
// result in compilation errors.
type UnsafeScrapyServer interface {
	mustEmbedUnimplementedScrapyServer()
}

func RegisterScrapyServer(s grpc.ServiceRegistrar, srv ScrapyServer) {
	s.RegisterService(&Scrapy_ServiceDesc, srv)
}

func _Scrapy_GetArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleScrapyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapyServer).GetArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wenworpc.scrapy.Scrapy/getArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapyServer).GetArticles(ctx, req.(*ArticleScrapyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Scrapy_ServiceDesc is the grpc.ServiceDesc for Scrapy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scrapy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wenworpc.scrapy.Scrapy",
	HandlerType: (*ScrapyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getArticles",
			Handler:    _Scrapy_GetArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wenworpc/scrapy.proto",
}
