package main

import "fmt"

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

func main() {
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println(aa, bb, cc)
}
