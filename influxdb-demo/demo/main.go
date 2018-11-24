/************************************************************
** @Description: demo
** @Author: haodaquan
** @Date:   2018-11-21 23:06
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-11-21 23:06
*************************************************************/
package main

import (
	"time"

	"fmt"

	"github.com/george518/go-study-notes/influxdb-demo/pool"
	client "github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB     = "pptest4"
	username = "george518	"
	password = "123456"
)

func main() {
	poolconfig := &pool.PoolConfig{
		InitialCap:  2,
		MaxCap:      10,
		Factory:     Factory,
		Close:       Close,
		Ping:        Ping,
		IdleTimeout: 5000 * time.Microsecond,
	}

	chanpools, err := pool.NewChannelPool(poolconfig)
	if err != nil {
		fmt.Println(err)
	}

	cl, err := chanpools.Get()
	if err != nil {
		fmt.Println(err)
	}

}

func Factory() (client.Client, error) {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://118.89.238.78:8086",
		Username: username,
		Password: password,
	})
	return c, err
}
func Close(c client.Client) error {
	return nil
}

func Ping(in interface{}) error {
	return nil
}
