package main

import (
	"fmt"
	"testing"
)

func Testtrys(t *testing.T) {
	res := acsc(2)
	if res != 2 {
		fmt.Println("cuowu")
	} else {
		fmt.Println("chenggong")

		t.Fatalf("%v,%v", 2, res)
		t.Logf("正确")
	}
}