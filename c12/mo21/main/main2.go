package main

import (
	"errors"
	"fmt"
	"os"
)

type circleQueue struct {
	maxsize int    //4
	array   [4]int //数组
	hand    int    //指向队首
	tail    int    //指向队尾
}

func (this *circleQueue) Addqueue() {

}

func (this *circleQueue) push(val int) (err error) {
	if this.isfull() {
		return errors.New("queue full")
	}

}

func (this *circleQueue) pop(val int) (val int, err error) {

}

func (this *circleQueue) isfull() bool {
	return ((this.tail+1)%this.maxsize == this.hand)

}

func (this *circleQueue) isempty() bool {
	return this.tail == this.hand

}

func (this *circleQueue) size() int {
	return (this.tail + this.maxsize - this.hand) % this.maxsize
}

func main() {
	var key string
	var val int

	for {
		fmt.Printf("1,输入add，表示获取数据到队列")
		fmt.Printf("2,输入get，表示从队列获取数据")
		fmt.Printf("3,输入show，表示显示队列")
		fmt.Printf("4,输入exit，表示退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入要入队列数")
			fmt.Scanln(&val)
			//err := queue.addqueue(val)
			//if err != nil {
			//	fmt.Printf("存在错误", err.Error())
			//} else {
			//	fmt.Printf("加入成功")
			//}
		case "get":
			fmt.Println("get")
		case "show":
			//queue.showqueue()
		case "exit":
			os.Exit(0)
		}
	}
}
