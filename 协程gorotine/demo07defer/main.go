package main

import "fmt"

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello ，world")
	}
}

func test() {
	//使用defer+recover解决
	defer func() {
		//捕获test（）抛出的异常 , 如果协程出现panic ，未捕获panic，就会出现异常。使用recover来捕获异常
		//即使协程发生异常，但主程序不受影响
		if err := recover(); err != nil {
			fmt.Printf("test()发生panic err=%v", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang"
}

func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}
}
