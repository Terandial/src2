package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var cmd, d string
	for {
		fmt.Println("每输入一个符号或数字记得按一下空格键")
		fmt.Println("请输入你的式子")
		fmt.Scan(&a)
		fmt.Scanln()
		fmt.Scan(&cmd)
		fmt.Scanln()
		fmt.Scan(&b)
		fmt.Scanln()

		if cmd == "+" {
			c = a + b

		}
		if cmd == "-" {
			c = a - b
		}
		if cmd == "*" {
			c = a * b

			if cmd == "/" {
				c = a / c
			}
			fmt.Println(c)
			fmt.Println("你想继续运算吗？ Y/N")
			fmt.Scanf("%s", &d)
			if d == "y" {
				fmt.Println("愿为您效劳")
				continue
			}
			if d == "n" {
				fmt.Println("感谢使用，期待下次见面")
				break
			}
		}

	}
}
