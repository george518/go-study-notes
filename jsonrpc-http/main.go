/************************************************************
** @Description: jsonrpc_http
** @Author: george hao
** @Date:   2018-08-29 10:10
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-29 10:10
*************************************************************/
package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(str string, reply *string) error {
	*reply = " http : " + str
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}

//启动：go run main.go
//$ curl localhost:1234/jsonrpc -X POST --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
//{"id":0,"result":" http : hello","error":null}
