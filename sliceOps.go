package main

import "fmt"

/**
切片的具体操作
*/

//打印Slice
func printSlice(s []int) {
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", s, len(s), cap(s))
}

func main() {

	//创建Slice
	var s []int
	var s1 = []int{2, 4, 6, 8}
	var s2 = make([]int, 16)
	var s3 = make([]int, 10, 20)

	//Slice添加元素
	for i := 1; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i)
	}

	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	//Slice复制元素
	copy(s2, s1)
	printSlice(s2)

	//Slice删除元素
	s2 = append(s2[:3], s2[4:]...) //删除第四个元素
	printSlice(s2)

	//Slice中得到第一个元素
	fmt.Printf("从Slice中得到第一个元素：%d\n", s2[0])

	//Slice中得到最后一个元素
	fmt.Printf("从Slice中得到最后一个元素：%d\n", s2[len(s2)-1])
}
