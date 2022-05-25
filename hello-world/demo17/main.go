package main

import "fmt"

func main() {
	res := sum(10, 20, 50, 10, 20, 20)
	fmt.Println("res=", res)
}

func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
