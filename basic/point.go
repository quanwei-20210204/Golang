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

	// 声明指针变量
	var ptr *int

	//指针变量 存储的地址
	ptr = &currentSecond

	fmt.Printf("the currentSecond address is: %x \n", &currentSecond)

	//指针变量 ptr 存储的地址
	fmt.Printf("the address contained in ptr is: %x \n", ptr)

	//使用指针 ptr 访问值
	fmt.Printf("the value using ptr access is: %d \n", *ptr)

	//当一个指针被定义后没有分配到任何变量时，它的值为 nil
	var ptrnil *int
	fmt.Printf("the ptrnil value is: %x \n", ptrnil)

	//空指针判断  yes   no
	//ptr 不是空指针  if(ptr != nil)
	//ptr 是空指针if(ptr == nil)

	if ptrnil == nil {
		fmt.Printf("the ptrnil value is: nil \n")
	}
}
