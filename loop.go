package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(number int) string {
	result := ""
	v := number
	for ; v > 0; v = v / 2 {
		lsb := v % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//func main() {
//	fmt.Println(
//		convertToBin(5),
//		convertToBin(11))
//
//	printFile("abc.txt")
//}
