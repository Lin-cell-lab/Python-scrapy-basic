package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	//1.获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v , type = %T , kind = %v", rVal, rVal, rVal)

	//2.转回interface{}
	iV := rVal.Interface()
	//类型断言
	n, ok := iV.(float64)
	if ok {
		fmt.Printf("n type = %T", n)
	}
}

func main() {
	var v float64 = 1.2
	reflect01(v)
}
