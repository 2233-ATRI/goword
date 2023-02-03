package main

import "fmt"

func main() {
	a := 0
	for i := 0; i < 10; i++ {
		fmt.Println((a + i) * (i + 1) / 2)
	}
}
