/************************************************************
** @Description: atomic
** @Author: george hao
** @Date:   2018-12-07 17:33
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-07 17:33
*************************************************************/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var count int32

func main() {

	wg.Add(1)
	start := time.Now().UnixNano()
	go func() {
		for i := 0; i < 1000000; i++ {
			atomic.AddInt32(&count, 1) //原子操作
		}
		wg.Done()
	}()

	for i := 0; i < 1000000; i++ {
		atomic.AddInt32(&count, 1)
	}

	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1000 / 1000)
	fmt.Println(count)

}
