package main

import (
	"fmt"
	"net"
	"os"

	// "github.com/robinxiang/socket_project/common"
	common "github.com/robinxiang/socket_project/common"
)

func handleConnectionFunc(conn net.Conn, serverConf common.Server_config) {
	var (
		this_server common.Server_config
		// server_start_msg string
		server_welcome string
		read_buffer    []byte
		write_buffer   []byte
		// conn             net.Conn
		read_msg_length int //read msg length
		// str_ip_port      string
	)

	//create read_buffer,write_buffer
	read_buffer = make([]byte, 1024)
	write_buffer = make([]byte, 1024)
	//send welcome message
	server_welcome = fmt.Sprintf("Welcome to SECRET socket...:%s\n", this_server.ServerIp)
	write_buffer = []byte(server_welcome)
	go func() {
		_, err := conn.Write(write_buffer)
		if err != nil {
			fmt.Printf("Send message error! %s", err)
			os.Exit(0)
		}
	}()

	defer conn.Close() //when exit close the net.conn
	//read data to the socket buffer
	read_msg_length, err := conn.Read(read_buffer)
	if err != nil {
		fmt.Printf("Read msg error!:\n%s", err)
	}
	fmt.Printf("got message(%d):%s", read_msg_length, string(read_buffer))
}

func main() {
	var (
		this_server      common.Server_config
		server_start_msg string
		// server_welcome   string

		conn net.Conn
		// read_msg_length  int //read msg length
		str_ip_port string
	)
	this_server = common.Server_config{
		ServerName:     "test server",
		ServerIp:       "127.0.0.1",
		ServerPort:     1122,
		ServerProtocol: "tcp",
	}
	for {
		server_start_msg = fmt.Sprintln("SCRET Socket Server is creating...")
		//make serverip:port format string
		str_ip_port = fmt.Sprintf("%s:%d", this_server.ServerIp, this_server.ServerPort)

		//create listenner
		listener, err := net.Listen(this_server.ServerProtocol, str_ip_port)

		//if create listener got err
		if err != nil {
			fmt.Printf("Can't listen on the %s:%d \n error quit!\n %s", this_server.ServerIp, this_server.ServerPort, err)
			os.Exit(0) // go err,exit the server program
		}

		//if listen has running
		fmt.Println(server_start_msg) // print the server start msg
		fmt.Printf("Server is listen on (%s)%s : %d\n", this_server.ServerProtocol, this_server.ServerIp, this_server.ServerPort)

		conn, err = listener.Accept()
		// if listener got error
		if err != nil {
			fmt.Printf("connection go error! %s\n error quit!", err)
			os.Exit(0)
		}
		// defer conn.Close() //when exit close the net.conn

		// start gorouteine to run handleConnectionFunc
		go handleConnectionFunc(conn, this_server)
	}
}
