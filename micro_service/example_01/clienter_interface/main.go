package main

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServer interface {
	Say(string, *string) error
}

type MyClient struct {
	ThisClient *rpc.Client
}

// init the client object
func InitClient(addr string) MyClient {
	client, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	return MyClient{ThisClient: client}
}

func (c *MyClient) Say(s string, reply *string) error {
	return c.ThisClient.Call("hello.Say", s, reply)
}

func main() {
	var test_result string
	test := InitClient("127.0.0.1:8181")
	err := test.Say("xiang tao", &test_result)
	if err != nil {
		fmt.Println("test.Say got error:", err)
	}
	fmt.Println(test_result)
}
