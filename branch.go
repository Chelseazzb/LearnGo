package main

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic("Invalid score")
		break
	case score < 60:
		g = "D"
		break
	case score < 80:
		g = "C"
		break
	case score < 90:
		g = "B"
		break
	case score <= 100:
		g = "A"
		break
	}
	return g
}

//func main() {
//	const file = "abc.txt"
//	contents, err := ioutil.ReadFile(file)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Printf("%s\n", contents)
//	}
//	fmt.Println(grade(61), grade(90))
//}
