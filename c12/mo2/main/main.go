package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D:/1111.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("err is %v", err)
		}
		fmt.Printf(str)
	}
	fmt.Println("文件读取结束")
}
