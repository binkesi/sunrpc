package srpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("TCP dial error:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	err = client.Call(DoubleServiceName+".DoubleNum", "42", &reply)
	if err != nil {
		log.Fatal("Service call error:", err)
	}
	fmt.Println(reply)
}
