package tests

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Devide(x, y int) float32 {
	if y == 0 {
		return 0.
	}
	return float32(x) / float32(y)
}

func UnitTest() {
	fmt.Println(Add(1, 2))
	fmt.Println(Devide(1, 2))
}
