package main

import (
	"fmt"
	"go-basics/calculator"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)

	sumResult := calculator.Sum(1, 2.1)
	fmt.Printf("a + b + offset = %g\n", sumResult)

}
