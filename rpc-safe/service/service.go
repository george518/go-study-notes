/************************************************************
** @Description: service
** @Author: george hao
** @Date:   2018-08-27 15:27
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-27 15:27
*************************************************************/
package main

import (
	"log"
	"net"
	"net/rpc"

	"fmt"

	"github.com/george518/go-study-notes/rpc-safe/service/base"
	. "github.com/george518/go-study-notes/rpc-safe/service/controller"
)

func main() {

	base.RegisterService(base.UserServiceName, new(User))
	base.RegisterService(base.ProductServiceName, new(Product))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		fmt.Println(conn)
		go rpc.ServeConn(conn)
	}
}
