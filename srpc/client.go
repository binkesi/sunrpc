package srpc

import (
	"log"
	"net/rpc"
)

type DoubleServiceClient struct {
	*rpc.Client
}

var _ DoubleServiceInterface = (*DoubleServiceClient)(nil)

func DialDoubleService(network, address string) (*DoubleServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		log.Fatal("RPC dial error:", err)
	}
	return &DoubleServiceClient{Client: client}, nil
}

func (client *DoubleServiceClient) DoubleNum(request string, reply *string) error {
	return client.Client.Call(DoubleServiceName+".DoubleNum", request, &reply)
}
