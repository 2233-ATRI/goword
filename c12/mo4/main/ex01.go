package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type charCount struct {
	charCount  int //英文个数
	Number     int //数字个数
	spacecount int //空格
	othercount int //其他字符
}

func main() {
	//打开文件进行reader访问读取，
	//	读取一行进行一行的处理
	fileName := "D:/1112.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	var catc charCount
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			//case v >= 'a':
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				catc.charCount++
			case v == ' ' || v == '\t':
				catc.Number++
			case v >= '0' && v <= '9':
				catc.spacecount++
			default:
				catc.othercount++
			}

		}
	}
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
		catc.charCount, catc.Number, catc.spacecount, catc.othercount)

}
