package main

import "fmt"

//斐波那契数列
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func main() {
	fmt.Println(fbn(7))
}
