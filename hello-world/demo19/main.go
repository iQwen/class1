package main

import "fmt"

func main() {
	//匿名函数
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	a := func(n3 int, n4 int) int {
		return n3 - n4
	}
	res2 := a(30, 40)
	fmt.Println("res2=", res2)

	//全局匿名函数
	var (
		Fun1 = func(n1 int, n2 int) int {
			return n1 * n2
		}
	)
	res4 := Fun1(20, 30)
	fmt.Println("res4=", res4)
}
