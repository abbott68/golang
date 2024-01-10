package main

import "fmt"

func main() {
	//声明一个数组
	var name [5]string

	//数组赋值
	name[0] = "abbott"
	name[1] = "test"
	name[2] = "t"
	//根据数组的索引打印值
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])
	// 声明数组的名称且指定大小，并赋值
	var num = [3]int{1, 2, 3}
	n := [5]int{1, 2, 4, 5, 6}
	fmt.Println(num[0])
	fmt.Println(n[0], n[1], n[2], n[3], n[4])
}
