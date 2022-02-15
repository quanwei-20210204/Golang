package main

import (
	"fmt"
	"time"
)

func main() {

	var currentSecond int = time.Now().Second()

	//变量是一种使用方便的占位符，用于引用计算机内存地址
	var a int = 10

	//Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	fmt.Printf("the address is: %x \n", &a)
	fmt.Printf("the address is: %x \n", &currentSecond)

	//了解了什么是内存地址和如何去访问它

	//一个指针变量指向了一个值的内存地址
	//定义指针变量。
	//为指针变量赋值。
	//访问指针变量中指向地址的值。

}
