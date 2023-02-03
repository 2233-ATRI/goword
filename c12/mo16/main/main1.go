package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rval := reflect.ValueOf(b)
	fmt.Println(rval.Kind())
	//其类型是指针类型，不是值类型
	//rval.SetInt(20)//不可以实现，
	//这里使用elem可以好的指针指向的内容
	rval.Elem().SetInt(20)
	//	rval.Elem() =
	//	num:= 20
	//	ptr *int = &num
	//	num2 := *ptr
}
func main() {
	var num int = 10
	reflect01(&num)
	//reflect01(num)
	//均报错
	//panic: reflect:
	//reflect.Value.SetInt using unaddressable value
	fmt.Println(num)
}
