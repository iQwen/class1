package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 6; j++ {
			if j == 4 {
				continue
			}
			fmt.Printf("j的值为%v \n", j)
		}
	}
}
