package srpc

import (
	"log"
	"strconv"
)

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
