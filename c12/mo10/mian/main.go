package main

import (
	"fmt"
	"time"
)

func tste() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello word")
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	go tste()
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang")
		time.Sleep(time.Second)
	}
	//num := runtime.NumCPU()
	//fmt.Println(num)

}

//
