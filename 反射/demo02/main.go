package main

import (
	"fmt"
)

func main() {
	//常量必须赋值
	var num int
	//必须初始化
	const tax int = 0
	// tax = 10 常量不能修改 只能修饰 bool 数值，string类型
	fmt.Println(num, tax)
	// const b = num / 2 num为变量不能使用
	const (
		//依次加1
		a    = iota
		b    = iota
		c, d = iota, iota //逐行递增
	)
	//大小写控制访问范围
	fmt.Println(a, b, c, d)
}
