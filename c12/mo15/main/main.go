package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	id   string
	Age  int
}

func (u user) reflectcallfunc() {
	fmt.Println("All.wu reflectcallect")
}

func main() {
	user := user{name: "wdbq", id: "1", Age: 12}
	Dofeladcin(user)
}

func Dofeladcin(u user) {
	getType := reflect.TypeOf(u)
	fmt.Println("get type is ", getType)
	//main.user
	fmt.Println("get type is ", getType.Name())
	//user
	getValue := reflect.ValueOf(u)
	fmt.Println("get all field is ", getValue)

	a := getValue.NumField()
	//该方法是对结构体内部的数据量进行统计
	fmt.Println(a)
	fmt.Println()
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		//value := getValue.Field(i).Interface()
		//博客中是这样描写的，但不可以实现，可能是这里的空接口办法有问题
		//无法返回从未导出的fie中获得的值
		//LD或方法
		fmt.Printf("%v: %v = %v\n", field.Name, field.Type, value)
	}
	fmt.Println()
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

//预期获得函数名和main.user，可能与结果获取有关
