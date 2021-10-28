package srpc

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func TestHttpServer(t *testing.T) {
	RegisterService(&Server{})
	http.HandleFunc("/jsonrpc", func(rw http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     rw,
			ReadCloser: r.Body,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}

func TestHttpClient(t *testing.T) {
	url := "http://127.0.0.1:1234/jsonrpc"
	var jsonstr = []byte(`{"method":"Server.DoubleNum", "params":["24"], "id":0}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonstr))
	if err != nil {
		log.Fatal("Form request error:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Response error:", err)
	}
	defer resp.Body.Close()
	fmt.Println("Response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read response error:", err)
	}
	fmt.Println("Response Body:", string(body))
}
