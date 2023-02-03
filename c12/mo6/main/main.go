package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Salt     float64
	skill    string
}

func sac() {
	str := "{\"name\":\"cnisacn\",\"Age\":500,\"birthday\":\"2022.1.2\", \"sal\":8000,\"skill\":\"sia\"}"
	//str :=   "{\"Name\":\"牛魔王~~~\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		//fmt.Println(err)
	}
	fmt.Printf("%v", monster)
}

// 演示将json字符串，反序列化成struct
func unmarshalStruct() {
	//说明str 在项目开发中，是通过网络传输获取到.. 或者是读取文件获取到
	str := "{\"Name\":\"牛魔王~~~\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	//定义一个Monster实例
	var monster Monster

	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v \n", monster, monster.Name)

}
func main() {
	sac()
	unmarshalStruct()
}
