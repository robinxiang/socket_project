package main

import (
	"fmt"
	"net"
	"os"

	common "github.com/robinxiang/socket_project/common"
)

func main() {
	var (
		this_server      common.Server_config
		server_start_msg string
		// server_welcome string
		read_buffer []byte
		// write_buffer []byte
		conn            net.Conn
		read_msg_length int //read msg length
	)
	this_server = common.Server_config{
		ServerName:     "test server",
		ServerIpPort:   "127.0.0.1:22",
		ServerProtocol: "tcp",
	}

	server_start_msg = fmt.Sprintf("SCRET Socket Server is creating...")

	//create listenner
	listener, err := net.Listen(this_server.ServerProtocol, this_server.ServerIpPort)

	//if create listener got err
	if err != nil {
		fmt.Printf("Can't listen on the %s:%d \n error quit!", this_server.ServerIp, this_server.ServerPort)
		os.Exit(0) // go err,exit the server program
	}

	//if listen has running
	fmt.Println(server_start_msg) // print the server start msg
	fmt.Printf("Server is listen on %s : %d\n", this_server.ServerProtocol, this_server.ServerIp)

	//create read_buffer,write_buffer
	read_buffer = make([]byte, 1024)
	// write_buffer = make([]byte, 1024)

	conn, err = listener.Accept()
	// defer conn.Close() //when exit close the net.conn
	// if listener got error
	if err != nil {
		fmt.Printf("connection go error! %s\n error quit!", err)
		os.Exit(0)
	}

	//read data to the socket buffer
	read_msg_length, err = conn.Read(read_buffer)
	fmt.Printf("got message(%d):%s", read_msg_length, string(read_buffer))

}
