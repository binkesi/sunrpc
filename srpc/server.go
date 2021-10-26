package srpc

import (
	"log"
	"net/rpc"
	"strconv"
)

const DoubleServiceName = "Server"

type DoubleServiceInterface interface {
	DoubleNum(request string, reply *string) error
}

func RegisterService(service DoubleServiceInterface) error {
	return rpc.RegisterName(DoubleServiceName, service)
}

type Server struct{}

func (server *Server) DoubleNum(request string, reply *string) error {
	if num, err := strconv.Atoi(request); err != nil {
		log.Println(err.Error())
		return err
	} else {
		*reply = strconv.Itoa(2 * num)
	}
	return nil
}
