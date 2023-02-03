package main

import (
	"fmt"
	"time"
)

func test() {
	var mymap map[int]string
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error", err)
		}
	}()
	mymap[0] = "asicnw"
}

func sellhello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,word")
	}
}

func main() {
	go sellhello()
	go test()

}
