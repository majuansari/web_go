package interceptor

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	"context"
)

func ClientAuthUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, exists := metadata.FromIncomingContext(ctx)
		if !exists {
			return nil, status.Errorf(codes.PermissionDenied, "%s is rejected as client id is missing.", info.FullMethod)
		}
		clientId := md.Get("client-id")
		secret := md.Get("client-secret")
		if clientId[0] != "10001" || secret[0] != "secret" {
			return nil, status.Errorf(codes.PermissionDenied, "%s is rejected as client id or secret is wrong.", info.FullMethod)
		}
		resp, err := handler(ctx, req)
		return resp, err
	}
}
