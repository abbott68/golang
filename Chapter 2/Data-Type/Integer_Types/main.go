package main

import (
	"fmt"
)

func main() {

	// 定义整数类型
	var a int = 1

	var b int = 22
	// 将b的值赋值给a ,将a变量的值重新赋值
	a = b

	fmt.Println(a)
	fmt.Println(b)
	//输出a+b的总和
	fmt.Println("c =", a+b)
}
