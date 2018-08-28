/************************************************************
** @Description: main
** @Author: george hao
** @Date:   2018-08-27 15:04
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-27 15:04
*************************************************************/
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "george", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
