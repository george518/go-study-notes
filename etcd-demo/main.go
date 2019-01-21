/************************************************************
** @Description: etcd_demo
** @Author: george hao
** @Date:   2019-01-11 12:28
** @Last Modified by:  george hao
** @Last Modified time: 2019-01-11 12:28
*************************************************************/
package main

import (
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:32773", "127.0.0.1:32772", "127.0.0.1:32769"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()
}
