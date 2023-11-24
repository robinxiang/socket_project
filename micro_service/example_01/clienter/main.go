package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

/*
RPC client ,practice the RPC protocol communication
*/
func main() {
	var (
		StrServerAddr   string
		StrServerPort   string
		StrServerNet    string
		StrServerString string
		strReply        string
		StrTest         string
	)
	StrServerAddr = "127.0.0.1"
	StrServerNet = "tcp"
	StrServerPort = "8181"
	StrServerString = fmt.Sprintf("%s:%s", StrServerAddr, StrServerPort)
	StrTest = "tao"

	// dail up to server
	// conn, err := rpc.Dial(StrServerNet, StrServerString)
	// practice the JSON RPC to call remote function
	conn, err := jsonrpc.Dial(StrServerNet, StrServerString)
	if err != nil {
		fmt.Println("rpc dial got error", err)
		return
	}
	defer conn.Close()

	// call the remote function
	err = conn.Call("people.Say", StrTest, &strReply)
	if err != nil {
		fmt.Println("rpc call got error", err)
		return
	}

	//print the RPC result from reply
	fmt.Println("RPC result is :", strReply)

}
