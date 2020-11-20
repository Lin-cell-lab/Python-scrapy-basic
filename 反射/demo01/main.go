package main

import (
	"fmt"
	"reflect"
)

//反射演示
func reflectTest(b interface{}) {
	//通过反射获取传入的type，kind，值
	//1.先获取relect。Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	//2.获取reflect.ValueOf

	eValue := reflect.ValueOf(b)

	n3 := 2 + eValue.Int()
	//n4 := eValue.Float()  call of reflect.Value.Float on int Value
	fmt.Println(n3)
	fmt.Println("eValue = ", eValue)
	fmt.Printf("eValue 类型为%T", eValue)

	//3.将rval转回interface{}
	iV := eValue.Interface()
	num2 := iV.(int)
	fmt.Println(num2)
}

//结构体反射
type students struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	//kind方法获取类别（）
	fmt.Printf("rtype kind = %v \n", rType.Kind())
	//type为类形 ， kind为类别 type和kind肯一样，肯不一样
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind = %v\n", rVal.Kind())
	//rVal 转回interfac{}
	iV := rVal.Interface()

	fmt.Printf("iv type = %T , iv value = %v\n", iV, iV)
	//断言取出需要的类型
	stu, ok := iV.(students)
	if ok {
		fmt.Println("name=", stu.Name)
	}
}
func main() {
	//演示对（基本数据类型，interfac{}，reflect。value）进行反射的基本操作
	var num int = 100
	reflectTest(num)

	//结构体案例
	stu := students{
		Name: "小米",
		Age:  20,
	}
	reflectTest02(stu)
}
