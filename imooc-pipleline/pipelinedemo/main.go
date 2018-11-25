/**********************************************
** @Des: This file...
** @Author: haodaquan
** @Date:   2018-01-22 00:19:06
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-01-22 22:55:46
***********************************************/
package main

import (
	"bufio"

	"fmt"
	"os"

	"github.com/george518/go-study-notes/imooc-pipleline/pipeline"
)

func main() {
	const filename = "small.in"
	const num = 64
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	p := pipeline.RandomSource(num)

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)

	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	i := 0
	for v := range p {
		fmt.Println(v)
		i++
		if i > 100 {
			break
		}
	}

	//p := pipeline.Merge(
	//	pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
	//	pipeline.InMemSort(pipeline.ArraySource(1, 9, 15, 2, 1,4)),
	//)
	//for v := range p {
	//	fmt.Println(v)
	//	// if num, ok := <-p; ok {
	//	// 	fmt.Println(num)
	//	// } else {
	//	// 	break
	//	// }
	//}

}
