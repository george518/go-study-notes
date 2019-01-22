/************************************************************
** @Description: operation
** @Author: george hao
** @Date:   2019-01-21 16:18
** @Last Modified by:  george hao
** @Last Modified time: 2019-01-21 16:18
*************************************************************/
package main

import (
	"time"

	"fmt"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
)

func main() {

	var (
		client *clientv3.Client
		err    error
		config clientv3.Config
		kv     clientv3.KV
		putOp  clientv3.Op
		getOp  clientv3.Op
		opResp clientv3.OpResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:32781", "127.0.0.1:32785", "127.0.0.1:32784"},
		DialTimeout: 5 * time.Second,
	}

	//建立连接
	client, err = clientv3.New(config)

	//建立kv
	kv = clientv3.NewKV(client)

	//创建Op
	putOp = clientv3.OpPut("/cron/jobs/job8", "123")

	//执行op
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	//获取验证
	fmt.Println("写入Version", opResp.Put().Header.Revision)

	//创建读op
	getOp = clientv3.OpGet("/cron/jobs/job8")
	//执行op
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("获取写入内容：", string(opResp.Get().Kvs[0].Value))

}
