package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "scinoa"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("scia")
	fmt.Printf("%v", str)
}
