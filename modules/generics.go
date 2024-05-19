package modules

import "fmt"

type customConstraints interface {
	// 独自の型: ~を付加すると含まれる
	~int | int16 | float32 | float64 | string
}

type NewInt int

func add[T customConstraints](x, y T) T {
	return x + y
}

func Generics() {
	fmt.Printf("%v\n", add(1, 2))
	fmt.Printf("%v\n", add(1.1, 2.1))
	fmt.Printf("%v\n", add("file", ".txt"))
	// bool型はcustomConstraintsに定義されていないためエラー
	// fmt.Printf("%v\n", add(true, false))
	var i1, i2 NewInt = 3, 4
	fmt.Printf("%v\n", add(i1, i2))
	fmt.Println("----------")
}
