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

func printForInfBreak(val int, thr int) {
	for {
		if val > thr {
			break
		}
		fmt.Printf("loop: %v 〜 %v\n", val, thr)
		val += 1
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
	// printForInf()

	e := 3
	printForInfBreak(e, 5)

	fmt.Println("----------")
}
