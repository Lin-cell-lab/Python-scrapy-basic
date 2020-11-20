package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)

	fmt.Println("kind = ", rVal.Kind())
	//将非地址转为地址相关值 ,设置对应的指针类型
	rVal.Elem().SetInt(20)

}
func main() {
	var num int = 21
	reflect01(&num)
	fmt.Println(num)
	// num3 := 9
	// prt * int = &num
	// num2 := *ptr

}
