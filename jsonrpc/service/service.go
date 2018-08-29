/************************************************************
** @Description: service
** @Author: george hao
** @Date:   2018-08-29 09:16
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-29 09:16
*************************************************************/
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(str string, reply *string) error {
	*reply = "jsonrpc :" + str
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		//注意ServerCodec是个方法，不是接口
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
