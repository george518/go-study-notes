/**********************************************
** @Des: 并发版helloworld
** @Author: haodaquan
** @Date:   2018-01-21 23:29:41
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-01-21 23:49:05
***********************************************/
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 50; i++ {
		go helloworld(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}

	// time.Sleep(1 * time.Millisecond)
}
func helloworld(i int, ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("hello world %d\n", i)
	}

}
