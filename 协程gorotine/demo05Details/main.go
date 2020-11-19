package main

import "fmt"

func main() {
	//管道可声明为可读或者可写默认情况下，可读也可写

	//1.只写
	var chan2 chan<- int

	chan2 = make(chan int, 3)

	chan2 <- 20
	//2.只读
	var chan3 <-chan int
	num := <-chan3
	fmt.Println(num)
	
}
