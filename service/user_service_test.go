package service

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"testing"
	"time"
	userpb "web/grpc/proto/user/go"
)

func BenchmarkCallTestGRPCAPI(b *testing.B) {
	opts := grpc.WithInsecure()
	con, err := grpc.Dial("127.0.0.1:50051", opts)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "web-api-client-us-east-1",
		"user-id", "some-test-user-id",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	client := userpb.NewUserDetailsServiceClient(con)
	request := &userpb.UserIdRequest{Id: "1"}
	for n := 0; n < b.N; n++ {
		client.UserDetailsService(ctx, request)
	}
}
func BenchmarkCallTestRestAPI(b *testing.B) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 10 * time.Second,
	}
	method := "POST"
	request := &userpb.UserIdRequest{Id: "1"}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(request)
	req, _ := http.NewRequest(method, "http://127.0.0.1:8000/user/test/1", payloadBuf)
	req.Header.Set("Accept", "application/json; charset=utf-8")

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		res, err := client.Do(req)

		ur := new(userpb.UserResponse)
		if err = json.NewDecoder(res.Body).Decode(ur); err != nil {
		}
	}
}
