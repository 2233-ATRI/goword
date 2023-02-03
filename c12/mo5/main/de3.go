package main

import (
	"encoding/json"
	"fmt"
)

//json

type Monster struct {
	Name     string `json:"monster_name"`
	Age      int
	Birthday string
	Salt     float64
	skill    string
}

func testStruct() {
	//演示结构体
	var ame Monster = Monster{
		Name:     "bew",
		Age:      12,
		Birthday: "2003.2.1",
		Salt:     8000.0,
		skill:    "wqwsnia",
	}
	//序列化
	data, err := json.Marshal(ame)
	if err != nil {
		fmt.Println("序列化错误 ", err)
	}
	fmt.Println(string(data))

}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "sabu"
	a["age"] = 13
	a["address"] = "cskalw"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化错误", err)
	}
	fmt.Println(string(data))
}

func testslice() {
	var slice []map[string]interface{}
	//切片类型的map
	var ma map[string]interface{}
	ma = make(map[string]interface{})
	//slice = make(map[string]interface{})
	ma["name"] = "jcnaw"
	ma["age"] = 1334
	ma["address"] = "beijing"
	slice = append(slice, ma)

	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	//slice = make(map[string]interface{})
	m1["name"] = "jcnaw"
	m1["age"] = 1334
	m1["address"] = "beijing"
	slice = append(slice, m1)
	data, err := json.Marshal(slice)

	if err != nil {
		fmt.Println("错误 ", err)

	}
	fmt.Println("序列化结果", string(data))
}

func main() {
	//对结构体，map，切片进行序列化
	testStruct()
	testMap()
	testslice()
}
