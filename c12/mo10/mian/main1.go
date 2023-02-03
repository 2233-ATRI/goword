package main

import "fmt"

func main() {
	var inchan chan interface{}
	inchan = make(chan interface{}, 10)

	fmt.Printf("%v %v\n", inchan, &inchan)
	inchan <- 10
	num := "saciqwd"
	inchan <- num
	fmt.Printf(" %v %v\n", len(inchan), cap(inchan))
	num1 := <-inchan
	fmt.Printf("%v", num1)
}
