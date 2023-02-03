package main

import (
	"fmt"
	"time"
)

func writedata(a chan interface{}) {
	for i := 0; i < 50; i++ {
		a <- i
		fmt.Println("write ", i)
	}
	close(a)
}

func readdate(a chan interface{}, b chan bool) {
	for {
		v, ok := <-a
		fmt.Println("read ", v)
		if !ok {
			break
		}
		fmt.Printf("%v", v)
	}
	b <- true
	close(b)
}
func main() {
	var a chan interface{}
	a = make(chan interface{}, 50)
	b := make(chan bool, 10)
	go readdate(a, b)
	go writedata(a)
	time.Sleep(time.Second * 10)
	defer close(b)
}

//只要向管道当做写入数据但没有进行读取，管道就会发生dead lock，
//原因是在管道容量内写入超过其内部体协程内的ch<-i
//主要原因是只写不读就会发生阻塞
//写管道和读管道的频率不同是不影响的
