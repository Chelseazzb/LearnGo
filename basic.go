package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 1
	bb = true
	cc = "abc"
)

func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("%d, %q\n", a, b)
}

func variableInitValue() {
	var a = 3
	var b = "123"
	var c = true
	fmt.Println(a, b, c)
}

func variableTypeDeduction() {
	var a, b, c = 3, "123", true
	fmt.Println(a, b, c)
}

func variableShorter() {
	a, b, c := 3, "123", true
	a = 5
	fmt.Println(a, b, c)
}

func euler() {
	//var a = 3 + 4i
	//fmt.Println(cmplx.Abs(a))

	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) //虚部不为0，因为complex类型的前后两部分都为float

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //类型转换是强制的
	fmt.Println(c)
}

func main() {
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println(aa, bb, cc)

	euler()

	triangle()
}
