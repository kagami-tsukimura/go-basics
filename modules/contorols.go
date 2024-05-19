package modules

import "fmt"

func Controls() {
	a := -1

	if a == 0 {
		fmt.Println("zero")
	} else if a > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}

	fmt.Println("----------")
}
