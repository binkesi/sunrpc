package revrpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"testing"
	"time"
)

func TestRevServer(t *testing.T) {
	rpc.RegisterName("ReverseServer", new(ReverseServer))
	for {
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
		conn.Close()
	}
}

func TestClient(t *testing.T) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("tcp listen error:", err)
	}
	clientChan := make(chan *rpc.Client)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("connection error:", err)
			}
			// pass rpc client to channel.
			clientChan <- rpc.NewClient(conn)
		}
	}()
	client := <-clientChan
	defer client.Close()
	var reply int
	args := [2]int{3, 4}
	err = client.Call("ReverseServer.Add", args, &reply)
	if err != nil {
		log.Fatal("call service error:", err)
	}
	fmt.Println(reply)
}
