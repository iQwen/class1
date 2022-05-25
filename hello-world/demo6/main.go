package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("啊啊啊啊啊啊啊啊啊啊")
	}

	j := 1
	for j <= 10 {
		fmt.Println("aaaaa")
		j++
	}

	k := 1
	for {
		if k <= 10 {
			fmt.Println("bbbbbbb")
		} else {
			break
		}
		k++
	}

	str := "hello,world!北京"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}
}
