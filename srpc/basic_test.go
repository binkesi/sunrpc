package srpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"reflect"
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
		go func() {
			ctx := jsonrpc.NewServerCodec(conn)
			//value := reflect.Indirect(reflect.ValueOf(ctx))
			vtype := reflect.TypeOf(reflect.ValueOf(ctx).Elem())
			for i := 0; i < vtype.NumField(); i++ {
				fmt.Printf("%s: %v\n", vtype.Field(i).Name, vtype.Field(i).Type)
			}
			rpc.ServeCodec(ctx)
		}()
	}
}

func TestJsonClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("TCP dial error:", err)
	}
	ctx := jsonrpc.NewClientCodec(conn)
	client := rpc.NewClientWithCodec(ctx)
	var reply string
	call := client.Go(DoubleServiceName+".DoubleNum", "42", &reply, make(chan *rpc.Call, 10))
	value := reflect.Indirect(reflect.ValueOf(call))
	fmt.Println(value)
	callrslt := <-call.Done
	err = callrslt.Error
	//err = client.Call(DoubleServiceName+".DoubleNum", "42", &reply)
	if err != nil {
		log.Fatal("Service call error:", err)
	}
	fmt.Println(reply)
}
