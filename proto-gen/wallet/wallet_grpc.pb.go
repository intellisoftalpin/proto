// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wallet

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

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletClient interface {
	DecodeTransaction(ctx context.Context, in *DecodeTransactionRequest, opts ...grpc.CallOption) (*DecodeTransactionResponse, error)
	SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	CheckTokenBalance(ctx context.Context, in *CheckTokenBalanceRequest, opts ...grpc.CallOption) (*Empty, error)
	GetAllTokens(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllTokensResponse, error)
	GetToken(ctx context.Context, in *TokenID, opts ...grpc.CallOption) (*GetTokenResponse, error)
	GetTokenPrice(ctx context.Context, in *TokenID, opts ...grpc.CallOption) (*GetTokenPriceResponse, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) DecodeTransaction(ctx context.Context, in *DecodeTransactionRequest, opts ...grpc.CallOption) (*DecodeTransactionResponse, error) {
	out := new(DecodeTransactionResponse)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/DecodeTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error) {
	out := new(SubmitTransactionResponse)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/SubmitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) CheckTokenBalance(ctx context.Context, in *CheckTokenBalanceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/CheckTokenBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetAllTokens(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllTokensResponse, error) {
	out := new(GetAllTokensResponse)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/GetAllTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetToken(ctx context.Context, in *TokenID, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetTokenPrice(ctx context.Context, in *TokenID, opts ...grpc.CallOption) (*GetTokenPriceResponse, error) {
	out := new(GetTokenPriceResponse)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/GetTokenPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
// All implementations must embed UnimplementedWalletServer
// for forward compatibility
type WalletServer interface {
	DecodeTransaction(context.Context, *DecodeTransactionRequest) (*DecodeTransactionResponse, error)
	SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	CheckTokenBalance(context.Context, *CheckTokenBalanceRequest) (*Empty, error)
	GetAllTokens(context.Context, *Empty) (*GetAllTokensResponse, error)
	GetToken(context.Context, *TokenID) (*GetTokenResponse, error)
	GetTokenPrice(context.Context, *TokenID) (*GetTokenPriceResponse, error)
	mustEmbedUnimplementedWalletServer()
}

// UnimplementedWalletServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (UnimplementedWalletServer) DecodeTransaction(context.Context, *DecodeTransactionRequest) (*DecodeTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeTransaction not implemented")
}
func (UnimplementedWalletServer) SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransaction not implemented")
}
func (UnimplementedWalletServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedWalletServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedWalletServer) CheckTokenBalance(context.Context, *CheckTokenBalanceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTokenBalance not implemented")
}
func (UnimplementedWalletServer) GetAllTokens(context.Context, *Empty) (*GetAllTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTokens not implemented")
}
func (UnimplementedWalletServer) GetToken(context.Context, *TokenID) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedWalletServer) GetTokenPrice(context.Context, *TokenID) (*GetTokenPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenPrice not implemented")
}
func (UnimplementedWalletServer) mustEmbedUnimplementedWalletServer() {}

// UnsafeWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServer will
// result in compilation errors.
type UnsafeWalletServer interface {
	mustEmbedUnimplementedWalletServer()
}

func RegisterWalletServer(s grpc.ServiceRegistrar, srv WalletServer) {
	s.RegisterService(&Wallet_ServiceDesc, srv)
}

func _Wallet_DecodeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).DecodeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/DecodeTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).DecodeTransaction(ctx, req.(*DecodeTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/SubmitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).SubmitTransaction(ctx, req.(*SubmitTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_CheckTokenBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CheckTokenBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/CheckTokenBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CheckTokenBalance(ctx, req.(*CheckTokenBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetAllTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAllTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/GetAllTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAllTokens(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetToken(ctx, req.(*TokenID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetTokenPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetTokenPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/GetTokenPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetTokenPrice(ctx, req.(*TokenID))
	}
	return interceptor(ctx, in, info, handler)
}

// Wallet_ServiceDesc is the grpc.ServiceDesc for Wallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DecodeTransaction",
			Handler:    _Wallet_DecodeTransaction_Handler,
		},
		{
			MethodName: "SubmitTransaction",
			Handler:    _Wallet_SubmitTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Wallet_GetTransaction_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _Wallet_CreateTransaction_Handler,
		},
		{
			MethodName: "CheckTokenBalance",
			Handler:    _Wallet_CheckTokenBalance_Handler,
		},
		{
			MethodName: "GetAllTokens",
			Handler:    _Wallet_GetAllTokens_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _Wallet_GetToken_Handler,
		},
		{
			MethodName: "GetTokenPrice",
			Handler:    _Wallet_GetTokenPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet/wallet.proto",
}