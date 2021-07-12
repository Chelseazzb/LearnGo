package main

import "fmt"

/**
学习go语言中数组的使用
go中一般不直接用数组，数组是值类型，值拷贝
*/

//打印数组的函数
func printArray(arr [5]int) {
	//只会暂时改变数组元素
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//打印数组的函数，指针会改变参数值
func printArray1(arr *[3]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {

	//数组有三种定义方式
	var arr1 = [5]int{}
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(arr1, arr2, arr3)

	//打印数组的两种方式
	fmt.Println("遍历range的arr3：")
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	fmt.Println("遍历range的arr3：")
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("打印arr1和arr3：")
	fmt.Println(arr1, arr3)

	fmt.Println("printArray1(arr2)")
	printArray1(&arr2)

}
