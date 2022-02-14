package main

import "fmt"

func main() {

	//
	var i, j, k int
	// 声明数组的同时快速初始化数组
	//var variable_name [SIZE] variable_type
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f \n", i, balance[i])

	}

	//声明不定长数组
	//var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var balance_var = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出每个数组元素的值 */
	for j = 0; j < 5; j++ {
		fmt.Printf("balance_var[%d] = %f \n", j, balance_var[j])
	}

	//将索引为 1 和 3 的元素初始化
	//指定下标来初始化元素 balance := [5]float32{1:2.0,3:7.0}
	var balance_choose = [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance_choose[%d] = %f \n", k, balance_choose[k])
	}

}