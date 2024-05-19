package modules

import "fmt"

type Os int

const (
	MAC Os = iota + 1
	WINDOWS
	LINUX
)

func Variables() {
	// variables
	// var
	var i int
	fmt.Println(i)
	var j int = 2
	fmt.Println(j)
	// 型推論
	var k = 4
	fmt.Println(k)
	fmt.Println("----------")

	// :=
	l := 1
	fmt.Println(l)
	fmt.Printf("i: %v %T\n", i, i)
	ui := uint16(6)
	// v: value
	// t: type
	fmt.Printf("ui: %v %T\n", ui, ui)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)
	fmt.Println("----------")

	f := 3.1
	fmt.Printf("f: %v %T\n", f, f)
	s := "hello"
	fmt.Printf("s: %v %T\n", s, s)
	b := true
	fmt.Printf("b: %v %T\n", b, b)
	fmt.Println("----------")

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v, title: %v\n", pi, title)
	fmt.Println("----------")

	// 型変換
	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Printf("x: %v, y: %v, z: %v\n", x, y, z)
	fmt.Println("----------")

	// 定数
	fmt.Printf("Mac: %v, Windows: %v, Linux: %v\n", MAC, WINDOWS, LINUX)
	fmt.Println("----------")

	// 変数の値を変更
	// l = 1 → 2
	fmt.Println(l)
	l = 2
	fmt.Println(l)
	l += 1
	fmt.Println(l)
	l *= 2
	fmt.Println(l)
	fmt.Println("----------")
}
