package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

// Clienter info struct
type Clienter struct {
	ClienterConn        net.Conn
	ClienterId          string
	ClienterIp          string
	ClienterName        string
	chanRead, chanWrite chan string
}

// make the rand ASCII string
func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
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
func handleConn(conn net.Conn, map_clienter map[string]Clienter, ChanBroadcast chan string) {
	var (
		strServerWelcome string
		strCastMsg       string
		strSendMsg       string
		intCmd           int64
	)

	strServerWelcome = fmt.Sprintf("welcome 2 b5t2!!!\n")
	strSendMsg = ""
	newClienter := Clienter{
		ClienterConn: conn,
		ClienterId:   conn.RemoteAddr().String(),
		ClienterIp:   conn.RemoteAddr().String(),
		ClienterName: generateRandomString(4),
		chanRead:     make(chan string, 1024),
		chanWrite:    make(chan string, 1024),
	}
	fmt.Printf("%s incoming\n", conn.RemoteAddr().String())
	//write new user online message to ChanBroadcast
	strCastMsg = fmt.Sprintf("[%s is join server from :%s]\n", newClienter.ClienterName, newClienter.ClienterId)
	ChanBroadcast <- strCastMsg
	//Add a new user client into maps
	map_clienter[newClienter.ClienterId] = newClienter

	// write welcome message to clients
	newClienter.chanWrite <- strServerWelcome

	// create gorouteine to send message to client
	go WriteBackClient(conn, &newClienter)
	//handle the client conn

	for {
		strReadData := make([]byte, 1024)
		int_read_length, err := conn.Read(strReadData)
		if err != nil {
			fmt.Printf("read data error:%s\n", err.Error())
			//write clinet offline message to ChanBroadcast
			strCastMsg = fmt.Sprintf("[%s is offline from :%s]\n", newClienter.ClienterName, newClienter.ClienterId)
			ChanBroadcast <- strCastMsg
			//delete the user client from maps
			delete(map_clienter, newClienter.ClienterId)
			return
		}
		//fmt.Printf("read data length:%d\n", int_read_length)
		//fmt.Printf("read data:%s\n", strReadData)
		strReadData = strReadData[:int_read_length]

		//if strReadData is \r\n or \n then continue
		if string(strReadData) == "\r\n" || string(strReadData) == "\n" {
			continue
		}

		strSendMsg, intCmd = HandleCmdList(&map_clienter, strings.Replace(strings.ToLower(string(strReadData)), "\n", "", -1))
		if intCmd == 0 {
			ChanBroadcast <- fmt.Sprintf("[%s]=>%s\n", newClienter.ClienterName, strSendMsg)
		} else {
			newClienter.chanWrite <- strSendMsg

		}

	}

}

func HandleCmdList(MapClient *map[string]Clienter, strMsg string) (string, int64) {
	var strResult string
	var intCmd int64
	strResult = ""
	intCmd = 0

	if strMsg == "who" {
		strResult = CmdWho(MapClient)
		intCmd = 1
	} else if strMsg[:7] == "rename" {

	}

	return strResult, intCmd
}

// handle the command "rename"
func CmdRename(Client *Clienter, strMsg string) string {
	//TODO
	var strResult string

	Client.ClienterName = strMsg
	strResult = fmt.Sprintf("Name was changed:%s", strMsg)
	return strResult

}

// handle the command "cmd"
func CmdWho(MapClient *map[string]Clienter) string {
	var str_who string
	for _, user := range *MapClient {
		str_who = str_who + fmt.Sprintf("%s\t%s\n", user.ClienterName, user.ClienterId)
		// str_who = str_who + user.ClienterName + "\n"
	}
	return str_who

}

// loop poling to handle the imcoming message from each user and put it in other user struct chan
func CareBroadcastChan(Broadcast chan string, map_client *map[string]Clienter) {
	//take care message for Broadcast by loop
	//start goroutein message show
	fmt.Println("Broadcast message process is start!")
	for {
		strMesasge := <-Broadcast
		for _, user := range *map_client {
			user.chanWrite <- strMesasge

		}
	}
}

// write back message to client
func WriteBackClient(con net.Conn, client *Clienter) {
	for data := range client.chanWrite {
		_, _ = con.Write([]byte(data))
	}
}

// the function to hanle the client read message
func (Clienter Clienter) handleReadMessage(conn net.Conn) {
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
	go CareBroadcastChan(ChanBroadcast, &map_clienter)

	for {
		// create conn
		conn, err := ServerListener.Accept()
		if err != nil {
			fmt.Println("Error on connection  accept:", err.Error())
			return
		}
		//get remote user net info

		go handleConn(conn, map_clienter, ChanBroadcast)
	}
}
