package main

import "fmt"

func main() {

	var totallevel int
	fmt.Printf("请输入金字塔得层数：")
	fmt.Scanf("%d", &totallevel)

	//i表示层数
	for i := 1; i <= totallevel; i++ {
		//在打印*前先打印空格
		for k := 1; k <= totallevel-i; k++ {
			fmt.Print(" ")
		}

		//j表示每层答应多少个
		for j := 1; j <= i*2-1; j++ {
			if j == 1 || j == i*2-1 || i == totallevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
