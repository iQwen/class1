package main

import "fmt"

func main() {
	count := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			fmt.Println(i)
			count++
			sum += i
		} else {
			fmt.Println(sum)
			fmt.Println(count)
		}
	}
}
