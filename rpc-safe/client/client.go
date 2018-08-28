/************************************************************
** @Description: client
** @Author: george hao
** @Date:   2018-08-27 17:01
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-27 17:01
*************************************************************/
package main

import (
	"fmt"
	"log"

	"github.com/george518/go-study-notes/rpc-safe/client/common"
)

func main() {
	client, err := common.DialService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.ProductCreate(" product create", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

	err = client.ProductGet(" product get ", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

	err = client.UserGet(" user get", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

	err = client.UserCreate(" user create", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
