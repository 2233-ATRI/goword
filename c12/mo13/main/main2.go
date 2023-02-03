package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 1.2345
	fmt.Println(num)
	fmt.Println()

	poerr := reflect.ValueOf(&num)
	nuwvaleica := poerr.Elem()
	fmt.Println("type of pointer :", nuwvaleica.Type())
	fmt.Println("settability of pointer :", nuwvaleica.CanSet())

	nuwvaleica.SetBool(true)
	fmt.Println(num)
	poerr = reflect.ValueOf(&num)
	nuwvaleica = poerr.Elem()
	fmt.Println("type of pointer :", nuwvaleica.Type())

}
