package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
	userpb "web/grpc/proto/user/go"
	appError "web/pkg/error"
)

func CallTestRestAPI(client *http.Client, url string, id string) (*userpb.UserResponse, error) {
	method := "POST"
	request := &userpb.UserIdRequest{Id: id}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(request)
	req, err := http.NewRequest(method, url, payloadBuf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Try to unmarshall into errorResponse
	if res.StatusCode != http.StatusOK {
		fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	ur := new(userpb.UserResponse)
	if err = json.NewDecoder(res.Body).Decode(ur); err != nil {
		return nil, err
	}

	return ur, nil
}

func CallTestGRPCAPI(ctx context.Context, con *grpc.ClientConn, id string) (*userpb.UserResponse, error) {
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "10001",
		"client-secret", "secret",
	)
	mdCtx := metadata.NewOutgoingContext(ctx, md)
	request := &userpb.UserIdRequest{Id: id}

	client := userpb.NewUserDetailsServiceClient(con)
	resp, err := client.UserDetailsService(mdCtx, request)
	//@todo handle grpc errors better
	if err != nil {
		errDetails := status.Convert(err)
		return nil, appError.NewRequestError(errDetails.Message(), http.StatusBadRequest, "grpc_error")
	}
	return resp, nil
}

//func CallTestGRPCAPI(ctx context.Context , address string, tracer trace.Tracer , id string) (*userpb.UserResponse, error) {
//	opts := grpc.WithInsecur	e()
//	con, err := grpc.Dial(address, opts,
//		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
//	)
//	if err != nil {
//		panic(err)
//	}
//	defer con.Close()
//
//	_, span := tracer.Start(ctx, "getUserGRPCClient", oteltrace.WithAttributes(attribute.String("id", id)))
//	defer span.End()
//
//	md := metadata.Pairs(
//		"timestamp", time.Now().Format(time.StampNano),
//		"client-id", "web-api-client-us-east-1",
//		"user-id", "some-test-user-id",
//	)
//	ctx = metadata.NewOutgoingContext(ctx, md)
//
//	client := userpb.NewUserDetailsServiceClient(con)
//	request := &userpb.UserIdRequest{Id: id}
//
//	resp, err := client.UserDetailsService(ctx, request)
//	return resp, err
//}
