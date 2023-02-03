package main

import "fmt"

type Valtype struct {
	row int //行
	col int //列
	avl int //值
}

func main() {
	//创建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑
	chessMap[2][3] = 2 //白
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
	}
	//转成稀疏数组
	//遍历之前生成的map 将其中不为特殊值的元素创建成为node节点
	//将其中放入对应的切片
	var sparseArr []Valtype

	valsac := Valtype{
		row: 11,
		col: 11,
		avl: 0,
	}
	sparseArr = append(sparseArr, valsac)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个值节点
				valsac := Valtype{
					row: i,
					col: j,
					avl: v2,
				}
				sparseArr = append(sparseArr, valsac)
			}
		}
	}
	fmt.Println("当前的稀疏数组是:::::")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.avl)
	}

	var chessMap2 [11][11]int
	for i, valsac := range sparseArr {
		if i != 0 {

			chessMap2[valsac.row][valsac.col] = valsac.avl
		}
	}
	fmt.Println("恢复之后的数据")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)

		}
		fmt.Println()
	}

}
