package tests

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
	// x, y := 3, 5
	// fmt.Println(Add(x, y))
	// fmt.Println(Devide(x, y))
	// fmt.Println(Devide(x, 0))
}
