package main

/**
Slice(切片)的底层机制
Slice是底层数组的一个视图，包括ptr（当前指针）
*/

//func main() {
//
//	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
//
//	//Slice切片
//	s1 := arr[:]
//	s2 := arr[2:6]
//	s3 := arr[2:]
//	s4 := arr[:7]
//
//	fmt.Println(s1)
//	fmt.Println(s2)
//	fmt.Println(s3)
//	fmt.Println(s4)
//
//	fmt.Println(arr)
//	s5 := arr[2:6]
//	s6 := s5[3:5]
//	fmt.Printf("s5=%v,len(s5)=%d,capacity(s5)=%d\n", s5, len(s5), cap(s5))
//
//	fmt.Printf("s6=%v,len(s6)=%d,capacity(s6)=%d\n", s6, len(s6), cap(s6))
//
//	//向slice中添加元素
//	s7 := append(s6, 100)
//	s8 := append(s7, 200)
//	fmt.Printf("s7=%v,len(s7)=%d,capacity(s7)=%d\n", s7, len(s7), cap(s7))
//	fmt.Printf("s8=%v,len(s8)=%d,capacity(s8)=%d\n", s8, len(s8), cap(s8))
//	fmt.Println(arr)
//
//}
