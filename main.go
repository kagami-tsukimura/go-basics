package main

import (
	"fmt"
	"go-basics/calculator"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// module, package
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)

	sumResultNumber := calculator.Sum(1, 2)
	sumResultFloat := calculator.Sum(1, 2.1)
	fmt.Printf("a + b + offset = %g\n", sumResultNumber)
	fmt.Printf("a + b + offset = %g\n", sumResultFloat)
	fmt.Printf("(a * b) + offset = %g\n", calculator.Multiply(1, 2))

	// variables

	// var
	var i int
	fmt.Println(i)
	var j int = 2
	fmt.Println(j)
	// 型推論
	var k = 4
	fmt.Println(k)

	// :=
	l := 1
	fmt.Println(l)
	fmt.Printf("i: %v %T\n", i, i)
	ui := uint16(6)
	fmt.Printf("ui: %v %T\n", ui, ui)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

}
