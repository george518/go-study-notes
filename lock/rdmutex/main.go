/************************************************************
** @Description: rdmutex
** @Author: george hao
** @Date:   2018-12-07 17:36
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-07 17:36
*************************************************************/
package main

import (
	"fmt"
	"sync"
	"time"
)

var rwLock sync.RWMutex //读写锁，读锁所有线程都可以同时用（除了写线程），但是同时写线程不能用写锁。用于读多写少。
var lock sync.Mutex
var w sync.WaitGroup
var count int

func main() {
	w.Add(1)
	start := time.Now().UnixNano()
	go func() {
		for i := 0; i < 1000; i++ {
			rwLock.Lock() //写锁
			//lock.Lock() //互斥锁
			count++
			time.Sleep(5 * time.Millisecond)
			//lock.Unlock()
			rwLock.Unlock()
		}
		w.Done()
	}()

	for i := 0; i < 16; i++ {
		w.Add(1)
		go func() {
			for i := 0; i < 5000; i++ {
				rwLock.RLock() //读锁
				//lock.Lock()
				time.Sleep(1 * time.Millisecond)
				//lock.Unlock()
				rwLock.RUnlock()
			}
			w.Done()
		}()
	}
	w.Wait()
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1000 / 1000)
	fmt.Println(count)
}
