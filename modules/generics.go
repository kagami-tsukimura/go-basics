package modules

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type customConstraints interface {
	// 独自の型: ~を付加すると含まれる
	~int | int16 | float32 | float64 | string
}

type NewInt int

func add[T customConstraints](x, y T) T {
	return x + y
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func sumValues[K int | string, V constraints.Float | constraints.Integer](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func Generics() {
	fmt.Printf("%v\n", add(1, 2))
	fmt.Printf("%v\n", add(1.1, 2.1))
	fmt.Printf("%v\n", add("file", ".txt"))
	fmt.Println("----------")

	// bool型はcustomConstraintsに定義されていないためエラー
	// fmt.Printf("%v\n", add(true, false))
	var i1, i2 NewInt = 3, 4
	fmt.Printf("%v\n", add(i1, i2))
	fmt.Printf("min value is %v\n", min(3, 4))
	fmt.Println("----------")

	m1 := map[string]uint{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	m2 := map[int]float32{
		1: 1.23,
		2: 4.56,
		3: 7.89,
	}

	fmt.Printf("%v\n", sumValues(m1))
	fmt.Printf("%v\n", sumValues(m2))
}
