package main

import (
	"bufio"
	"channel-study/pipeline"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//网络版
	p := createNetworkPipeline("small.in", 512, 4)
	writeToFile(p, "small_net.out")
	printFile("small_net.out")

	//p := createPipeline("small.in",512,4)
	//writeToFile(p ,"small.out")
	//printFile("small.out")

	// 4-8-18
	// 1-5-7
	// 8 3-10-13

	//p := createPipeline("large.in",8000000,8)
	//writeToFile(p ,"large.out")
	//printFile("large.out")
}
func printFile(s string) {
	file, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count > 20 {
			break
		}
	}
}
func writeToFile(p <-chan int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func createPipeline(fileName string, fileSize, chunkCount int) <-chan int {
	var sortResults []<-chan int
	chunkSize := fileSize / chunkCount //存在除不尽的情况
	pipeline.Init()                    //初始化开始时间
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)

		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResults...)
}

func createNetworkPipeline(fileName string, fileSize, chunkCount int) <-chan int {
	var sortAddr []string
	chunkSize := fileSize / chunkCount //存在除不尽的情况
	pipeline.Init()                    //初始化开始时间
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)

		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...)
}
