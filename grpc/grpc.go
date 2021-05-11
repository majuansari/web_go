package grpc

import "google.golang.org/grpc"

func InitGrpC(address string, opts ...grpc.DialOption) (*grpc.ClientConn, func()) {
	con, err := grpc.Dial(address, opts...)
	if err != nil {
		panic(err)
	}
	cleanUp := func() {
		con.Close()
	}
	return con, cleanUp
}
