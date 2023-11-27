package main

import (
	"context"
	"fmt"

	"github.com/robinxiang/socket_project/micro_service/example_protobuf/server/invoicer/github.com/robinxiang/socket_project/micro_service/example_protobuf/server/invoicer"
	"google.golang.org/grpc"
)

func main() {
	// grpc.Dial to connect server
	grpcConn, err := grpc.Dial("127.0.0.1:8089", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dail to server error:", err)
	}
	defer grpcConn.Close()

	//create grpc client
	thisClient := invoicer.NewInvoicerClient(grpcConn)

	//create the message
	var CreateRequest invoicer.CreateRequest
	// CreateRequest = new(invoicer.CreateRequest)
	CreateRequest.Amount = &invoicer.Amount{
		Amount:   20,
		Currency: "hello",
	}
	CreateRequest.Amount.Currency = "hello"
	CreateRequest.From = "hello"
	CreateRequest.To = "hello"

	// call remote rpc function
	i, err := thisClient.Create(context.TODO(), &CreateRequest)
	if err != nil {
		fmt.Println("Create error:", err)
	}

	fmt.Println(i)

}
