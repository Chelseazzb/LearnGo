package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// go里面的可变参数列表
func sum(numbers ...int) int {
	var i int
	for i := range numbers {
		i += i
	}
	return i
}

// go的运算处理函数
func op(a, b int, c string) (int, error) {
	switch c {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		//出错处理，但是会直接终止执行
		//panic("Unsupported operation : " + c)
		return 0, fmt.Errorf("Unsupported operation : %s", c)
	}
}

//求除和余的函数
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//函数式编程，接受op函数作为参数
func apply(op func(int, int) int, a, b int) int {
	//拿到函数的名字
	p := reflect.ValueOf(op).Pointer()
	funcName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args : (%d,%d)\n", funcName, a, b)
	return op(a, b)
}

//求方函数
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func swap(a, b *int) {
	*b, *a = *a, *b

}

func main() {
	fmt.Println(op(1, 3, ";"))
	fmt.Println(div(13, 3))
	fmt.Println(apply(pow, 2, 3))
	// 匿名函数
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 2, 3))

	//指针的使用
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
