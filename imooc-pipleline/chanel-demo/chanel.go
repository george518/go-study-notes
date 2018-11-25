/************************************************************
** @Description: chanel_demo
** @Author: haodaquan
** @Date:   2018-03-12 23:45
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-12 23:45
*************************************************************/
package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for v := range c {
		fmt.Printf("recived %d,worked:%c \n", id, v)
	}

	//for {
	//	if v, ok := <-c; ok {
	//		fmt.Printf("recived %d,worked:%c \n", id, v)
	//	} else {
	//		fmt.Println("yes")
	//	}
	//}

}

func createWork(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func bufferChannel() {
	c := make(chan int, 2)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Second)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Second)
}

func chanDomo() {
	var channels [10]chan<- int

	for i := 0; i < 9; i++ {
		channels[i] = createWork(i)
	}

	for i := 0; i < 9; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 9; i++ {
		channels[i] <- 'A' + i
	}

	//c = make(chan int)
	//go work(1, c)
	//go func() {
	//	c <- 1
	//	c <- 2
	//}()

	//
	//for i := 0; i < 9; i++ {
	//	go func() {
	//		fmt.Println(<-channels[i])
	//	}()
	//}

	time.Sleep(time.Second)
}
func main() {
	//chanDomo()
	//bufferChannel()
	channelClose()
}
