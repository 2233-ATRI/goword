package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type nodeval struct {
	row int
	col int
	val interface{}
}

//文件读取转成原始数据
func ReadData(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()
	bfrd := bufio.NewReader(file)
	var index = 0
	var arr [][]int
	for {
		line, err := bfrd.ReadBytes('\n')
		if err != nil {
			break
		}
		index++
		temp := strings.Split(string(line), " ")
		row, _ := strconv.Atoi(temp[0])
		col, _ := strconv.Atoi(temp[1])
		value, _ := strconv.Atoi(temp[2])
		if index == 1 {
			for i := 0; i < row; i++ {
				var arr_temp []int
				for j := 0; j < col; j++ {
					arr_temp = append(arr_temp, value)
				}
				arr = append(arr, arr_temp)
			}
		}
		if index != 1 {
			arr[row][col] = value
		}
	}
	// 打印数据
	fmt.Println()
	fmt.Println("从磁盘读取后的数据")
	fmt.Println()
	for _, v := range arr {
		for _, v1 := range v {
			fmt.Printf("%d\t", v1)
		}
		fmt.Println()
	}
}

func main() {

	var chessmap [11][11]int
	chessmap[1][2] = 1
	chessmap[2][3] = 2

	// 看看原始数据
	for _, v := range chessmap {
		for _, v1 := range v {
			fmt.Printf("%d\t", v1)
		}
		fmt.Println()
	}

	// 转成稀疏数据
	var sparseArr []nodeval
	// 数据规模
	sparseArr = append(sparseArr, nodeval{
		row: 11,
		col: 11,
		val: 0,
	})
	//稀疏数组
	for row, val := range chessmap {
		for col, val1 := range val {
			if val1 != 0 {
				sparseArr = append(sparseArr, nodeval{
					row: row,
					col: col,
					val: val1,
				})
			}
		}
	}
	// 稀疏数组存盘
	filepath := "F:/1233.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, node := range sparseArr {
		str := fmt.Sprintf("%d %d %d \n", node.row, node.col, node.val)
		writer.WriteString(str)
	}
	writer.Flush()

	// 稀疏数据从磁盘读取转换成原始数据
	ReadData(filepath)
}

