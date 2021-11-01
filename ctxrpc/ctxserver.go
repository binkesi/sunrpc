package ctxrpc

import "net"

type CtxServer struct {
	conn net.Conn
}

func (server *CtxServer) Hello(request string, reply *string) error {
	*reply = "hello: " + request + server.conn.LocalAddr().String()
	return nil
}
