/************************************************************
** @Description: done
** @Author: haodaquan
** @Date:   2018-03-14 23:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-14 23:43
*************************************************************/
package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWork(i int, worker worker) {
	for {
		fmt.Printf("id：%d,worker:%c \n", i, <-worker.in)
		worker.done()
	}
}

func createWork(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(i, w)
	return w
}

func channelDone() {
	var w [10]worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		w[i] = createWork(i, &wg)
	}

	//for n, ws := range w {
	//	go doWork(n, ws.in, &wg)
	//}

	wg.Add(20)
	for n, ws := range w {
		ws.in <- n + 'a'
	}

	for n, ws := range w {
		ws.in <- n + 'A'
	}

	wg.Wait()
}
func main() {

	channelDone()
}
