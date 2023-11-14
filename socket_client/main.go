package main

import (
	"fmt"
	"net"

	common "github.com/robinxiang/socket_project/common"
)

func main() {
	var (
		this_server       common.Server_config
		str_server_port   string
		conn              net.Conn
		byte_SendDate     []byte
		byte_RecvDate     []byte
		str_SendDate      string
		str_RecvDate      string
		int_MessageLength int
	)

	this_server = common.Server_config{
		ServerName:     "test server",
		ServerIp:       "127.0.0.1",
		ServerPort:     1122,
		ServerProtocol: "tcp",
	}

	byte_SendDate = make([]byte, 1024)
	byte_RecvDate = make([]byte, 1024)

	str_server_port = fmt.Sprintf("%s:%d", this_server.ServerIp, this_server.ServerPort)
	fmt.Println("SECRET socket client is conneting...")
	conn, err := net.Dial("tcp", str_server_port)
	if err != nil {
		fmt.Println("SECRET socket client is error:", err)
		return
	}
	defer conn.Close() //when exit ,close the connection

	//show message connected
	fmt.Printf("Connected to server:%s port:%d", this_server.ServerIp, this_server.ServerPort)
	str_SendDate = "I'm robin"
	byte_SendDate = []byte(str_SendDate)
	int_MessageLength, err = conn.Write(byte_SendDate)
	if err != nil {
		fmt.Println("SECRET socket client is error:", err)
		return
	}

	int_MessageLength, err = conn.Read(byte_RecvDate)

	str_RecvDate = string(byte_RecvDate[:int_MessageLength])
	fmt.Println("Received message from server:", str_RecvDate)
}
