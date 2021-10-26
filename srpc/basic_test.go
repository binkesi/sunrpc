package srpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"testing"
)

// This is just a test file to try functionality

func TestRpc(t *testing.T) {
	fmt.Println("This is a test.")
}

func TestDouble(t *testing.T) {
	RegisterService(&Server{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Net listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Connection error:", err)
		}
		go rpc.ServeConn(conn)
	}
}

func TestClient(t *testing.T) {
	client, err := DialDoubleService("tcp", ":1234")
	if err != nil {
		log.Fatal("RPC dial error:", err)
	}
	var reply string
	err = client.DoubleNum("23", &reply)
	if err != nil {
		log.Fatal("Service error:", err)
	}
	fmt.Println(reply)
}
