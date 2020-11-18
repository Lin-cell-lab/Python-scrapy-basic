package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 29
	intChan <- 30
	//1.关闭管道后，不能写入数据，但是仍然可以读取
	close(intChan)
	n1 := <-intChan
	fmt.Println("n1 = ", n1)
	//2.遍历时，channel未关闭，则出现deadlock错误，如果channel已经关闭，则会正常遍历数据，遍历完成后，就自动退出。
	//for range遍历

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
