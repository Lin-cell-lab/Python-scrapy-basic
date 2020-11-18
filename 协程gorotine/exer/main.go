package main

import (
	"fmt"
)

type cat struct {
	Name string
	Age  int
}

func main() {
	//定义一个存放任意数据类型的管道3个数据
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "小李"
	cat1 := cat{Name: "小米", Age: 12}
	allChan <- cat1

	//希望获得第三个
	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat=%v,newCat=%T\n", newCat, newCat)
	//写法错误，使用类型断言
	// fmt.Println("newCat.Name=%v", newCat.Name)
	a := newCat.(cat)
	fmt.Println(a.Name)

}
