/************************************************************
** @Description: 续租的问题
** @Author: george hao
** @Date:   2019-01-11 15:32
** @Last Modified by:  george hao
** @Last Modified time: 2019-01-11 15:32
*************************************************************/
package main

import (
	"time"

	"fmt"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
)

func main() {

	var client *clientv3.Client
	var config clientv3.Config
	var err error
	var lease clientv3.Lease
	var leaseGrantResp *clientv3.LeaseGrantResponse
	var leaseId clientv3.LeaseID
	var keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
	var keepResp *clientv3.LeaseKeepAliveResponse
	var kv clientv3.KV
	var putResp *clientv3.PutResponse
	var getResp *clientv3.GetResponse

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:32781", "127.0.0.1:32785", "127.0.0.1:32784"},
		DialTimeout: 5 * time.Second,
	}

	//建立一个客户端
	client, err = clientv3.New(config)

	if err != nil {
		fmt.Println(err)
		return
	}

	//申请一个lease
	lease = clientv3.NewLease(client)

	//申请一个10s租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	//拿到租约ID
	leaseId = leaseGrantResp.ID

	// 5秒后会取消自动续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil {
		fmt.Println("自动续租")
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepResp == nil {
					fmt.Println("租约已经失效了")
					goto END
				} else {
					// 每秒会续租一次, 所以就会受到一次应答
					fmt.Println("收到自动续租应答", keepResp.ID)
				}
			}
		}

	END:
	}()

	// 获得kv API子集
	kv = clientv3.NewKV(client)

	// Put一个KV, 让它与租约关联起来, 从而实现10秒后自动过期
	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入成功:", putResp.Header.Revision)

	// 定时的看一下key过期了没有
	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期:", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}

}
