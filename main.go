package main

import (
	"fmt"
	"go-basics/calculator"
	"os"

	"github.com/joho/godotenv"
)

const SECRET = "abc"

type Os int

const (
	MAC Os = iota + 1
	WINDOWS
	LINUX
)

var (
	i int
	s string
	b bool
)

func modulePackage() {
	// module, package
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println("----------")

	sumResultNumber := calculator.Sum(1, 2)
	sumResultFloat := calculator.Sum(1, 2.1)
	fmt.Printf("a + b + offset = %g\n", sumResultNumber)
	fmt.Printf("a + b + offset = %g\n", sumResultFloat)
	fmt.Printf("(a * b) + offset = %g\n", calculator.Multiply(1, 2))
	fmt.Println("----------")
}

func variables() {
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

func pointers() {
	var ui1 uint16
	// pointer: メモリ内の1byte番地
	// 先頭アドレス
	fmt.Printf("memory address og ui1: %p\n", &ui1)
	var ui2 uint16
	// pointer: メモリ内の1byte番地
	// 先頭アドレス
	fmt.Printf("memory address og ui2: %p\n", &ui2)

	// pointer変数: *<型>
	// NOTE: pointer変数は型の宣言が必須
	// NOTE: pointerで変数の先頭番地が分かり、型で長さが分かる
	var p1 *uint16
	// NIL
	fmt.Printf("value of p1: %v\n", p1)
	// ui1のアドレス情報をp1に代入
	fmt.Println("Assign ui1 address information to p1")
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)

	fmt.Println("----------")
}

func main() {

	// modulePackage()
	// variables()
	pointers()

}
