package main

import (
	"context"
	"log"
	"net"

	"github.com/robinxiang/socket_project/micro_service/example_protobuf/server/invoicer/github.com/robinxiang/socket_project/micro_service/example_protobuf/server/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test pdf"),
		Docx: []byte("test Docx"),
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal("create listener error%s", err)
	}
	serverRegister := grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegister, service)
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatal("create serve error:%s", err)

	}
}
