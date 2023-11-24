package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Person struct {
	FirstName string
	Age       int
	Msg       string
}

// bind the method to service
func (p *Person) Say(LastName string, resp *string) error {
	*resp = p.FirstName + " " + LastName + " " + p.Msg
	return nil
}

func main() {
	//show create service info
	fmt.Println("Starting RPC server....")
	//regist service to RPC server
	err := rpc.RegisterName("people", new(Person))
	if err != nil {
		fmt.Println("registry service error:", err)
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
	rpc.ServeConn(conn)

}
