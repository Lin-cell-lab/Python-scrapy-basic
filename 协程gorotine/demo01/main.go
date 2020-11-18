package main

import (
	"fmt"
)

type cat struct {
	Name string
	Age  int
}

func main() {
	//管道的使用  ,引用类型
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intchan的地址为%p\n", &intChan)
	fmt.Printf("intchan的值为%v\n", intChan)

	//写入数据
	intChan <- 10
	num := 12
	intChan <- num
	intChan <- 21

	//查看长度和cap
	fmt.Printf("channel 的长度和容量：length=%v, cap =%v\n", len(intChan), cap(intChan))

	//管道中读取数据
	var n2 int
	n2 = <-intChan
	fmt.Println("n2 = ", n2)
	fmt.Printf("channel 的长度和容量：length=%v, cap =%v", len(intChan), cap(intChan))

	n3 := <-intChan
	n4 := <-intChan
	fmt.Println("n3 = ", n3)
	fmt.Println("n4 = ", n4)
	//如果在没有使用协程的情况下，如果管道的数据全部取出，再取就会报错，deadlock
	// n5 := <-intChan
	// fmt.Println("n5 = ", n5)

	//存放map
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["小米"] = "介休"
	m1["凳子"] = "广安"

	m2 := make(map[string]string, 20)
	m2["小米1"] = "介休1"
	m2["凳子1"] = "广安1"
	mapChan <- m1
	mapChan <- m2
	m3 := <-mapChan
	fmt.Println(m3)

	//存放结构体
	var catChan chan cat
	catChan = make(chan cat, 20)
	cat1 := cat{Name: "小米", Age: 20}
	cat2 := cat{Name: "小邓", Age: 21}
	catChan <- cat1
	catChan <- cat2
	cat3 := <-catChan
	fmt.Printf("cat3=%v,cat3类型为：%T", cat3, cat3)
}
