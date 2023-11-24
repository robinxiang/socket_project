package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// use interface encapsulate JSON RPS server

type HelloServer interface {
	Say(string, *string) error
}

// standardize parameter incoming through interface
func RegisterHelloService(s HelloServer) error {
	return rpc.RegisterName("hello", s)
}

type MessageService struct {
}

func (m *MessageService) Say(msg string, reply *string) error {
	*reply = "hello,i got " + msg
	return nil
}

func main() {
	err := RegisterHelloService(new(MessageService))
	if err != nil {
		fmt.Println("RegisterHelloService got err:", err)
	}

	//create RPC listener
	listener, err := net.Listen("tcp", "0.0.0.0:8181")
	if err != nil {
		fmt.Println("create listen error:", err)
	}
	defer listener.Close()

	//get connection
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("conn accept error:", err)
	} else {
		fmt.Printf("RPC server started!!!")
	}
	defer conn.Close()

	//bind rpc server
	// rpc.ServeConn(conn)

	// practice JSON RPC server
	jsonrpc.ServeConn(conn)

}
