package main

import "fmt"

func main() {

	totalsum := 0.0
	passcount := 0
	for j := 1; j <= 3; j++ {
		sum := 0.0
		for i := 1; i <= 5; i++ {
			var score float64
			fmt.Printf("请输入第%d班的第%d个学生的成绩 \n", j, i)
			fmt.Scanln(&score)
			sum += score
			if passcount >= 60 {
				passcount++
			}
		}

		fmt.Printf("第%d个班级的平均分是%v \n", j, sum/5)
		totalsum += sum
	}

	fmt.Printf("所有班的平均分为%v", totalsum/3)
	fmt.Printf("及格人数为%v \n", passcount)
}
