package main

import "fmt"

func main() {

	var a int = 1
	var b int = 2
	fmt.Println(bool(a == b)) //判断条件a和b是否相等
	fmt.Println(bool(a != b)) //判断条件a and b 不相等
	fmt.Println(bool(a > b))  //判断条件 a 是否大于 b
	fmt.Println(bool(a < b))  //判断条件 a 是否小于b
	fmt.Println()
}
