package main

import (
	"context"
	"fmt"
	"net"

	"github.com/robinxiang/socket_project/micro_service/example_grpc_consul/pb/github.com/robinxiang/socket_project/micro_service/example_grpc_consul/pb"

	"google.golang.org/grpc"
)

type Childen struct {
}

// bind class
func (c *Childen) SayHello(ctx context.Context, in *pb.Person, opts ...grpc.CallOption) (*pb.Person, error) {
	in.Name = "Hello " + in.Name
	return in, nil
}

func main() {
	// init grpc object
	grpcServer := grpc.NewServer()

	// register service
	pb.RegisterHelloServer(grpcServer, new(Childen))

	// setup listener
	listenner, err := net.Listen("tcp", "111.47.170.131:8090")
	if err != nil {
		fmt.Println("setup listener error:", err)
	}
	defer listenner.Close()

	// start server

}
