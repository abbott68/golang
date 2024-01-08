package main

import (
	"fmt"
)

func main() {

	const a = 1 //声明一个产量

	const b = 2

	const (
		c    = b
		age  = 19
		name = "abbott"
	)
	fmt.Println(a + c)
	fmt.Println("怎么称呼您:" + name)
	fmt.Println("您的年龄是:", age)

}
