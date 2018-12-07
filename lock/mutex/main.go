/************************************************************
** @Description: mutex
** @Author: george hao
** @Date:   2018-12-07 17:35
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-07 17:35
*************************************************************/
package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex  //互斥锁
var w sync.WaitGroup //等待子线程退出
var count int

func main() {
	start := time.Now().UnixNano()
	w.Add(1) //相当于标记起一个子线程
	go func() {
		for i := 0; i < 1000000; i++ {
			lock.Lock()
			count++
			lock.Unlock()
		}
		w.Done() //相当于标记关闭一个子线程
	}()

	for i := 0; i < 1000000; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}

	w.Wait()
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1000 / 1000)
	fmt.Println(count)

}
