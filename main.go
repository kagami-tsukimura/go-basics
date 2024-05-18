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

	sumResult := fmt.Sprintf("a + b + offset = %f", calculator.Sum(1, 2))

	fmt.Println(sumResult)
}
