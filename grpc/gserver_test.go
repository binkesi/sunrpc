package grpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"testing"
	"time"

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
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal("stream error:", err)
	}
	//create a seperate go routine to send message.
	go func() {
		for i := 0; i < 8; i++ {
			if err = stream.Send(&String{Value: "hi"}); err != nil {
				log.Fatal("send error:", err)
			}
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()
	//loop to recieve message on main routine.
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("send channel closed")
				break
			}
			log.Fatal("recieve error:", err)
		}
		fmt.Println(reply.GetValue())
	}
}
