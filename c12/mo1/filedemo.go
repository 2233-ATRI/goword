package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:/1111.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	fmt.Printf("file = %v", file)
	err = file.Close()
	if err != nil {
		fmt.Println("close file err", err)
	}
}
