package main

import "fmt"

func main() {
	n1 := 10
	test(n1)
	fmt.Println("main n1=", n1)
}

func test(n1 int) int {
	n1 = n1 + 1
	fmt.Println("test n1=", n1)
	return n1
}
