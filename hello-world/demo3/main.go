package main

import (
	"fmt"
)

func main() {
	var I int = 10
	fmt.Println(&I)
	var ptr *int = &I
	fmt.Println(ptr)
	fmt.Printf("ptr 的地址是%v\n", &ptr)
	fmt.Printf("ptr 指向的值是%v\n", *ptr)
	*ptr = 8
	fmt.Printf("ptr的值改为了%v", *ptr)
}
