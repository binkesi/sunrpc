package kvsvc

import (
	"log"
	"net"
	"net/rpc"
	"testing"
	"time"
)

func TestKvService(t *testing.T) {
	rpc.RegisterName("KvStoreService", NewKvStoreService())
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("net listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("connection error:", err)
		}
		go rpc.ServeConn(conn)
	}
}

func TestWatchClient(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("rpc dial error:", err)
	}
	go func() {
		var keyChanged string
		err := client.Call("KvStoreService.Watch", 5, &keyChanged)
		if err != nil {
			log.Fatal("call service error:", err)
		}
		log.Println("watch:", keyChanged)
	}()
	time.Sleep(2 * time.Second)
	err = client.Call("KvStoreService.Set", [2]string{"name", "sungn"}, new(struct{}))
	if err != nil {
		log.Fatal("call service error:", err)
	} else {
		log.Println("set key")
	}
	time.Sleep(10 * time.Second)
}
