package main

import (
	"fmt"
	"time"
)

//放入
func putNum(intChan chan int) {
	for i := 1; i <= 10000; i++ {
		intChan <- i
		// fmt.Printf("放入%v\n", i)
	}
	close(intChan)
}

// 开启4个协程，从intChan取Chan出数据
func primeChanF(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break //退出管道for循环
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//num非素数
				flag = false
				break
			}
			flag = true
			if flag {
				//放入primeChan
				primeChan <- num
			}
		}
	}
	fmt.Println("有一个primeChan因为取不到数据而退出！")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	//退出标识符
	exitChan := make(chan bool, 8)

	//时间统计
	start := time.Now().Unix()
	//1.开启协程，放入1-8000个整数
	go putNum(intChan)

	//2.开启4个协程，从intChan取出数据，判断是否未素数，如果是就放入primeChan
	for i := 0; i < 8; i++ {
		go primeChanF(intChan, primeChan, exitChan)
	}

	//3.处理主线程
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		//取出四个结果
		end := time.Now().Unix()
		fmt.Println("程序使用时间为：", end-start)
		close(primeChan)
	}()
	//4.遍历primeChan取出数据
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		// fmt.Printf("素数%d\n", res)
	}
	// fmt.Println("主线程退出!")
}
