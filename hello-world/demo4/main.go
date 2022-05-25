package main

import (
	"fmt"
	"math"
)

func main() {
	//var name string
	//var age byte
	//var sal float32
	//var isPass bool
	//fmt.Println("请输入姓名")
	//fmt.Scanln(&name)
	//
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//
	//fmt.Println("请输入薪水")
	//fmt.Scanln(&sal)
	//
	//fmt.Println("请输入是否通过考试(false or true)")
	//fmt.Scanln(&isPass)
	//
	//fmt.Printf("名字是%v \n年龄是%v \n薪水是%v \n通过了考试%v \n", name, age, sal, isPass)
	//fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试(使用空格隔开)")
	//fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	//fmt.Printf("名字是%v \n年龄是%v \n薪水是%v \n通过了考试%v \n", name, age, sal, isPass)

	var a, b, c float64
	fmt.Println("请输入a,b,c")
	fmt.Scanf("%f %f %f", &a, &b, &c)
	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		x2 := (-b - math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=%v x2=%v", x1, x2)
	} else if m == 0 {
		x := (-b + math.Sqrt(m)) / (2 * a)
		fmt.Printf("x=%v", x)
	} else {
		fmt.Println("方程无解")
	}
}
