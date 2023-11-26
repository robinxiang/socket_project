package main

import (
	"log"
	"net"

	"github.com/robinxiang/socket_project/micro_service/example_protobuf/server/invoicer/github.com/robinxiang/socket_project/micro_service/example_protobuf/server/invoicer"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal("create listener error%s", err)
	}
	serverRegister := grpc.NewServer
	invoicer.RegisterInvoicerServer()
}
