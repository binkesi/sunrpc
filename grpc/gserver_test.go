package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
)

func TestGrpcServer(t *testing.T) {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("net listen error:", err)
	}
	grpcServer.Serve(listener)
}

func TestGrpcClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc dial error:", err)
	}
	defer conn.Close()
	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "hello"})
	if err != nil {
		log.Fatal("call service error:", err)
	}
	fmt.Println(reply.GetValue())
}
