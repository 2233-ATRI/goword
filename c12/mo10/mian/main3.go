package main

import "fmt"

func main() {
	var a chan int
	a = make(chan int, 100)
	for i := 0; i < 50; i++ {
		a <- i
	}
	close(a)
	fmt.Printf(" %v %v\n", len(a), cap(a))
	//遍历时没有关闭管道的话会发生死锁，在关闭之后会正常遍历，然后退出
	for i := 0; i < 100; i++ {

		for v := range a {
			fmt.Printf("%v", v)
			//管道只能获得一个值所以不可以使用两个值进行接收
		}
	}
}
