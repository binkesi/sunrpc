package srpc

import (
	"io"
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
