/************************************************************
** @Description: _select
** @Author: haodaquan
** @Date:   2018-03-16 23:50
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-16 23:50
*************************************************************/
package main

import (
	"fmt"
	"math/rand"

	"time"
)

func worker(id int, c chan int) {
	for v := range c {
		time.Sleep(time.Second)
		fmt.Printf("recived %d,worked:%d \n", id, v)
	}
}

func createWork(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			//time.Sleep(2 * time.Second)
			c <- i
			i++
		}
	}()
	return c
}

func main() {
	c1, c2 := generator(), generator()
	w := createWork(0)
	var values []int

	tm := time.After(10 * time.Second)
	tt := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out")
		case <-tt:
			fmt.Println("queue len is ", len(values))
		case <-tm:
			fmt.Println("Bye")
			return
		}
	}

}
