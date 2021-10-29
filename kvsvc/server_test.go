package kvsvc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"testing"
	"time"
)

func TestKvService(t *testing.T) {
	kvstore := NewKvStoreService()
	rpc.Register(kvstore)
	listener, err := net.Listen("tcp", ":2424")
	if err != nil {
		log.Fatal("net listen error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("connection error:", err)
	}
	for {
		go func() {
			rpc.ServeConn(conn)
		}()
	}
}

func TestWatchClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:2424")
	if err != nil {
		log.Fatal("connection error:", err)
	}
	client := rpc.NewClient(conn)
	go func() {
		var keyChanged string
		err := client.Call("KvStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal("call service error:", err)
		}
		fmt.Println("watch:", keyChanged)
	}()
	err = client.Call("KvStoreService.Set", [2]string{"name", "sungn"}, new(struct{}))
	if err != nil {
		log.Fatal("call service error:", err)
	}
	time.Sleep(3 * time.Second)
}
