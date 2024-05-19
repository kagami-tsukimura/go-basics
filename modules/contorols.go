package modules

import "fmt"

func printIf(a int) {
	if a == 0 {
		fmt.Println("zero")
	} else if a > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}
}

func Controls() {
	a := -1
	printIf(a)

	fmt.Println("----------")
}
