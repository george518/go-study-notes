/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2018-01-22 00:13:13
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-01-22 22:51:57
***********************************************/
package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {

	out := make(chan int, 1024) //有缓存通道
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		fmt.Println(" Read done :", time.Now().Sub(startTime))

		//排序
		sort.Ints(a)
		fmt.Println(" Sort done :", time.Now().Sub(startTime))

		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2

		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)

		fmt.Println(" Merge done :", time.Now().Sub(startTime))
	}()
	return out
}

func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}

			if err != nil || (bytesRead == -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[:m]...))
}
