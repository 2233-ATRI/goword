package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	//mymap := map[int]int
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁lock
	lock sync.Mutex
)

// test 函数就是计算 n!, 让将这个结果放入到 myMap
func saca(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i

	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()

}
func main() {
	//cpunum := runtime.NumCPU()
	//fmt.Println(cpunum)
	//runtime.GOMAXPROCS(cpunum - 1)
	//fmt.Println(cpunum)

	for i := 1; i <= 20; i++ {
		go saca(i)
	}
	time.Sleep(time.Second * 5)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
	lock.Unlock()
}
