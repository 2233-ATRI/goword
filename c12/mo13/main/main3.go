package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c4
		d
	)
	var num int
	const asc int = 123
	fmt.Println(asc, num)
	// const c = getval()
}
func getval() {

}

//type和kind可能相同也有可能不同，var num int = 10 这样是相同的
//但对于结构体而言其中的type是包名.结构体名，kind是struct
