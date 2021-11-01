package ctxrpc

import (
	"fmt"
	"log"
	"net"
)

type CtxServer struct {
	conn    net.Conn
	isLogin bool
}

func (server *CtxServer) Login(request string, reply *string) error {
	if request != "username:password" {
		server.isLogin = false
		return fmt.Errorf("login failed")
	}
	log.Println("login ok")
	server.isLogin = true
	return nil
}

func (server *CtxServer) Hello(request string, reply *string) error {
	if server.isLogin == false {
		return fmt.Errorf("please login first")
	}
	*reply = "hello: " + request + server.conn.LocalAddr().String()
	return nil
}
