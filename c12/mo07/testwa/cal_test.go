package main

import (
	"fmt"
	"testing"
)

func TestAddupper(t *testing.T) {
	//调用
	res := addupper(10)
	if res != 55 {
		fmt.Println("cuowu")
	} else {
		fmt.Println("chenggong")

		t.Fatalf("%v,%v", 55, res)
		t.Logf("正确")
	}
}
