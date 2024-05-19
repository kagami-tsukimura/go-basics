package modules

import (
	"fmt"
	"time"
)

func printIf(val int) {
	if val == 0 {
		fmt.Println("zero")
	} else if val > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}
	fmt.Println("----------")
}

func printFor(val int) {
	for i := 0; i < val; i++ {
		fmt.Println(i)
	}
	fmt.Println("----------")
}

func printForInf() {
	for {
		fmt.Println("infinite loop")
		time.Sleep(2 * time.Second)
	}
}

func Controls() {
	a := 0
	printIf(a)
	b := 1
	printIf(b)
	c := -1
	printIf(c)

	d := 10
	printFor(d)
	printForInf()

	fmt.Println("----------")
}
