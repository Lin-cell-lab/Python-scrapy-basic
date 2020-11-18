package main

import (
	"fmt"
)

//如果只是向管道写入数据而没有读取，则会出现阻塞
//写入数据到管道
func writeDate(intChan chan int) {
	for i := 1; i < 100; i++ {
		intChan <- i
		fmt.Printf("写入数据%v\n", i)
	}
	// time.Sleep(time.Second)
	close(intChan)
}

func readDate(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Sleep) 异步处理
		fmt.Printf("读取数据%v\n", v)
	}
	//读取完数据 ， 任务完成
	exitChan <- true
	close(exitChan)
}

// 如果没有缓冲区，单纯的往其中放入元素立马就会进入阻塞状态，必须有其他的线程从其中取走元素。通俗的讲要有一个线程不断的取这个管道的元素，才能往其中放入元素。它就像一个窄窄的门框，进去就得出来。
// 而有一个缓冲区的管道想一段地道，放入的元素不会马上进入阻塞状态，只有第二个准备进入而第一个还没有进入的情况下才会阻塞。
func main() {
	intChan := make(chan int, 100)
	exitChan := make(chan bool, 1)
	go writeDate(intChan)
	go readDate(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
