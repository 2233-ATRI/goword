package main

import "fmt"

func addupper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	n := 10
	fmt.Println(addupper(n))
}
