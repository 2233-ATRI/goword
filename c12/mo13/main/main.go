package main

import (
	"fmt"
	"reflect"
)

func repctTest(b interface{}) {
	tType := reflect.TypeOf(b)
	fmt.Print("rptype", tType)
}

func main() {
	var num int = 100
	repctTest(num)
}
