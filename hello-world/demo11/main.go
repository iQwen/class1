package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	for i := 100; i < 1000; i++ {
		a = (i / 100)
		b = (i / 10 % 10)
		c = (i % 10)
		if i == a*a*a+b*b*b+c*c*c {
			fmt.Println(i)
		}
	}
}
