package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "D:/1111.txt"
	contest, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read err is %v", err)
	}
	fmt.Printf("%v", string(contest))
}
