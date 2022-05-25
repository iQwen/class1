package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var num int
	var count int = 0
	fmt.Println("请输入一个您想生成的随机数")
	fmt.Scanf("%d", &num)
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if n == num {
			break
		}

		fmt.Printf("生成%d一共用了%v次\n", num, count)
	}
}
