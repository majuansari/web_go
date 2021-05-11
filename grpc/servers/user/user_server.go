package user

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"web/app"
	userpb "web/grpc/proto/user/go"
)

type Server struct {
	*app.App
}

func (u *Server) UserDetailsService(ctx context.Context, req *userpb.UserIdRequest) (*userpb.UserResponse, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "grpc.user.user_details")
	defer span.End()
	id := req.Id
	fmt.Println("id", id)
	return &userpb.UserResponse{
		Name:  "Maju",
		City:  "TDPA",
		Phone: "1234444",
	}, nil
}
