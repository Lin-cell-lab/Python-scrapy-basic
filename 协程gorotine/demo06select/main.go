package main

import (
	"fmt"
)

func main() {
	//整数管道
	intChan := make(chan int, 5)
	for i := 0; i < 5; i++ {
		intChan <- i
	}

	//字符串管道
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "xioami" + fmt.Sprintf("%d", i)
	}

	// 传统方法在遍历的时候，不关闭管道会阻塞而导致deadlock
	//问题，使用select解决管道关闭的时机问题
	label := false
	for {
		select {
		case v := <-intChan:
			fmt.Printf("读取intChan数据%v\n", v)
		case v := <-strChan:
			fmt.Printf("读取strChan数据%v\n", v)
		default:
			fmt.Println("未读取到数据，不玩了，程序员写入逻辑！")
			label = true
		}
		if label {
			break
		}
	}

}
