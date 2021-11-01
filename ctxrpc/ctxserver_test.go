package ctxrpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"testing"
)

func TestCtxServer(t *testing.T) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("tcp listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go func() {
			defer conn.Close()
			p := rpc.NewServer()
			p.RegisterName("HelloServer", &CtxServer{conn: conn})
			p.ServeConn(conn)
		}()
	}
}

func TestCtxClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net dial error:", err)
	}
	request := "sungn2"
	var reply string
	client := rpc.NewClient(conn)
	err = client.Call("HelloServer.Login", "username:password", &reply)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Call("HelloServer.Hello", request, &reply)
	if err != nil {
		log.Fatal("call service error:", err)
	}
	fmt.Println(reply)
}
