package main

//数组模拟环形队列
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
	//最后一个数据位不可以存元素，不然无法进行环形化
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxsize
	return
}

func (this *circleQueue) pop() (val int, err error) {
	if this.isempty() {
		return 0, errors.New("queue empty")
	}
	//取出hand是指向队首，包含队首元素
	val = this.array[this.hand]
	this.hand = (this.hand + 1) % this.maxsize
	return
}

func (this *circleQueue) listqueue() {
	size := this.size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	temphead := this.hand
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", temphead, this.array[temphead])
		temphead = (temphead + 1) % this.maxsize
	}

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
	queue := &circleQueue{
		maxsize: 5,
		hand:    0,
		tail:    0,
	}
	var key string
	var val int

	for {
		fmt.Printf("1,输入add，表示获取数据到队列\n")
		fmt.Printf("2,输入get，表示从队列获取数据\n")
		fmt.Printf("3,输入show，表示显示队列\n")
		fmt.Printf("4,输入exit，表示退出\n")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入要入队列数")
			fmt.Scanln(&val)
			err := queue.push(val)
			if err != nil {
				fmt.Printf("存在错误", err.Error())
			} else {
				fmt.Printf("加入成功")
			}
		case "get":
			fmt.Println("get")
			val, err := queue.pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("取出了一个数为 =", val)
			}
		case "show":
			queue.listqueue()
		case "exit":
			os.Exit(0)
		}
	}
}
