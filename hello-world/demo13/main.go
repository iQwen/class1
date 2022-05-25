package main

import "fmt"

func main() {

	var name string
	var pwd string
	var chance = 3

label1:
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入您的名字")
		fmt.Scanln(&name)
		fmt.Println("请输入您的密码")
		fmt.Scanln(&pwd)

		if name == "cqupt" && pwd == "888888" {
			fmt.Println("登陆成功！")
			break
		} else {
			chance--
			if chance == 0 {
				break label1
			}
			fmt.Printf("您还有%v次机会\n", chance)
		}
	}

	if chance == 0 {
		fmt.Println("次数已用完，账号已冻结。")
	}

}
