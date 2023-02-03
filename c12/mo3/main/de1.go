package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := "D:/1112.txt"
	filepath, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file error ", err)
		return
	}
	defer filepath.Close()
	//准备写入
	str := "hello word\n"
	writer := bufio.NewWriter(filepath)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}
