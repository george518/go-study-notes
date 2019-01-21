/************************************************************
** @Description: demo1
** @Author: george hao
** @Date:   2019-01-11 12:52
** @Last Modified by:  george hao
** @Last Modified time: 2019-01-11 12:52
*************************************************************/
package main

import (
	"context"
	"time"

	"fmt"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	var config clientv3.Config
	var client *clientv3.Client
	var err error
	var kv clientv3.KV

	//var ctx context.Context

	var respose *clientv3.PutResponse
	var getRes *clientv3.GetResponse
	var delRes *clientv3.DeleteResponse

	//配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:32781", "127.0.0.1:32785", "127.0.0.1:32784"},
		DialTimeout: 5 * time.Second,
	}

	//创建链接
	client, err = clientv3.New(config)
	defer client.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)

	//新增kv
	respose, err = kv.Put(context.TODO(), "/test/kvs", "da", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(respose)

	//查询
	getRes, err = kv.Get(context.TODO(), "/test/pi")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(getRes)

	//更新
	respose, err = kv.Put(context.TODO(), "/test/pi", "hao")
	fmt.Println(respose)

	//再查询
	getRes, err = kv.Get(context.TODO(), "/test/pi")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(getRes)

	//删除
	delRes, err = kv.Delete(context.TODO(), "/test/pi", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(delRes)

	//再查询
	getRes, err = kv.Get(context.TODO(), "/test/pi")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(getRes)
	//wc := client.Watch(context.Background(), "/test/", clientv3.WithPrefix(), clientv3.WithPrevKV())
	//for v := range wc {
	//	if v.Err() != nil {
	//		panic(err)
	//	}
	//	for _, e := range v.Events {
	//		fmt.Printf("type:%v\n kv:%v  prevKey:%v  ", e.Type, e.Kv, e.PrevKv)
	//	}
	//}

}
