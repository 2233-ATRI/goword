package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxsize int
	array   [5]int //数组模拟队列
	front   int    //指向队列队首
	rear    int    //指向队列队尾
}

// 添加数据到队列
func (this *Queue) addqueue(val int) (err error) {
	//判断队列是否满
	if this.rear == this.maxsize-1 {
		return errors.New("queue full")

	}
	this.rear++
	this.array[this.rear] = val
	return
}

func (this *Queue) getqueue() (val int, err error) {
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

// 显示队列
func (this *Queue) showqueue() {
	for i := this.front + 1; i < this.rear; i++ {
		fmt.Printf("arrary[%d] = %d\t", i, this.array[i])
	}
}

func main() {
	queue := &Queue{
		maxsize: 5,
		front:   -1,
		rear:    -1,
	}
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
			err := queue.addqueue(val)
			if err != nil {
				fmt.Printf("存在错误", err.Error())
			} else {
				fmt.Printf("加入成功")
			}
		case "get":
			fmt.Println("get")
		case "show":
			queue.showqueue()
		case "exit":
			os.Exit(0)
		}
	}
}
