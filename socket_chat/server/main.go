package main

import (
	"fmt"
	"net"
)

// Clienter info struct
type Clienter struct {
	ClienterConn        net.Conn
	ClienterId          string
	ClienterIp          string
	ClienterName        string
	chanRead, chanWrite chan string
}

// transform []byte to string list by length rule
func byte2String(data []byte, int_length int) []string {
	var str_list []string
	for i := 0; i < len(data); i += int_length {
		str_list = append(str_list, string(data[i:i+int_length]))
	}
	return str_list
}

// the function to handle the conn
func (Clienter Clienter) handleConn(conn net.Conn) {
	fmt.Println("USER_ID:%s", Clienter.ClienterId)
	conn.Write([]byte("Hello"))
	return
}

// loop poling to handle the imcoming message from each user and put it in other user struct chan
func CareBroadcastChan(Broadcast chan string) {
	//take care message for Broadcast by loop
	for {

	}
}

func main() {
	var (
		ServerIp, ServerNetwork string
		ServerPort              int
		ServerListener          net.Listener //the server listener,tcp
		intMsgLength            int
		strServerWelcome        string
		map_clienter            map[string]Clienter
		ChanBroadcast           chan string
	)
	// server welcome message
	strServerWelcome = `
	bbbbbbbb
	b::::::b            555555555555555555          tttt           222222222222222
	b::::::b            5::::::::::::::::5       ttt:::t          2:::::::::::::::22
	b::::::b            5::::::::::::::::5       t:::::t          2::::::222222:::::2
	 b:::::b            5:::::555555555555       t:::::t          2222222     2:::::2
	 b:::::bbbbbbbbb    5:::::5            ttttttt:::::ttttttt                2:::::2
	 b::::::::::::::bb  5:::::5            t:::::::::::::::::t                2:::::2
	 b::::::::::::::::b 5:::::5555555555   t:::::::::::::::::t             2222::::2
	 b:::::bbbbb:::::::b5:::::::::::::::5  tttttt:::::::tttttt        22222::::::22
	 b:::::b    b::::::b555555555555:::::5       t:::::t            22::::::::222
	 b:::::b     b:::::b            5:::::5      t:::::t           2:::::22222
	 b:::::b     b:::::b            5:::::5      t:::::t          2:::::2
	 b:::::b     b:::::b5555555     5:::::5      t:::::t    tttttt2:::::2
	 b:::::bbbbbb::::::b5::::::55555::::::5      t::::::tttt:::::t2:::::2       222222
	 b::::::::::::::::b  55:::::::::::::55       tt::::::::::::::t2::::::2222222:::::2
	 b:::::::::::::::b     55:::::::::55           tt:::::::::::tt2::::::::::::::::::2
	 bbbbbbbbbbbbbbbb        555555555               ttttttttttt  22222222222222222222	
	`
	// server config
	ServerIp = "0.0.0.0"
	ServerNetwork = "tcp"
	ServerPort = 2121
	intMsgLength = 1024

	// Broadcast chanel made
	ChanBroadcast = make(chan string)
	//client map
	map_clienter = make(map[string]Clienter, 10)

	// show start message to server screem
	fmt.Printf("Socket chat server starting....\n")

	// create server to listen
	ServerListener, err := net.Listen(ServerNetwork, fmt.Sprintf("%s:%d", ServerIp, ServerPort))
	if err != nil {
		fmt.Println("Error create listening:", err.Error())
		return
	}
	fmt.Printf("Server is listen on %s:%d\n", ServerIp, ServerPort)
	fmt.Println(strServerWelcome, intMsgLength)
	for {
		// create conn
		conn, err := ServerListener.Accept()
		if err != nil {
			fmt.Println("Error on connection  accept:", err.Error())
			return
		}
		//get remote user net info
		newClienter := Clienter{
			ClienterConn: conn,
			ClienterId:   fmt.Sprintf("%s:%d", conn.RemoteAddr().String(), conn.Port().String),
			ClienterIp:   conn.RemoteAddr().String(),
		}
		fmt.Printf("%s incoming\n", conn.RemoteAddr().String())
		//Add a new user client into maps
		map_clienter[newClienter.ClienterId] = newClienter
		//write new user online message to ChanBroadcast
		ChanBroadcast <- fmt.Sprintf("%s is join server from :%s", newClienter.ClienterName, newClienter.ClienterId)
		conn.Write([]byte(strServerWelcome))
		//handle the client conn
		go map_clienter[conn.RemoteAddr().String()].handleConn(conn)
	}
}
