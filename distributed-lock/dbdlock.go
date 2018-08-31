/************************************************************
** @Description: distributed_lock
** @Author: george hao
** @Date:   2018-08-31 09:21
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-31 09:21
*************************************************************/
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("无锁情况下：")
	nolock()
	fmt.Println("有锁情况下：")
	haslock()
}

//无锁
var nolockCount int

//有锁
var l sync.Mutex
var lockCount int

func nolock() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nolockCount++
		}()
	}
	wg.Wait()
	fmt.Println(nolockCount)
}

func haslock() {
	var wgl sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wgl.Add(1)
		go func() {
			defer wgl.Done()
			//进程内加锁
			l.Lock()
			lockCount++
			l.Unlock()
		}()
	}
	wgl.Wait()
	fmt.Println(lockCount)
}
