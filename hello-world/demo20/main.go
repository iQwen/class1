package main

import (
	"fmt"
	"strings"
)

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {

	fmt.Printf("%p\n", &suffix)

	return func(name string) string {
		fmt.Printf("%p\n", &suffix)
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

	f1 := makeSuffix(".jpg")
	fmt.Println("文件处理后名字为:", f1("winter"))
	fmt.Println("文件处理后名字为:", f1("111.jpg"))
}
