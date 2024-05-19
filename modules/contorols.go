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
			fmt.Println("loop done!")
			break
		}
		fmt.Printf("loop: %v 〜 %v\n", val, thr)
		val += 1
		time.Sleep(2 * time.Second)
	}
}

func printForSwitch(val int) {
	// loop: switch内のbreakでfor loopを抜けられるように命名
loop:
	for i := 0; i < val; i++ {
		switch i {
		case 2:
			continue
		case 3:
			// NOTE: switchだけ抜ける時はcontinue
			continue
		case 8:
			// NOTE: breakのみではswitchだけ抜けるためループ継続
			// loop自体を抜ける
			break loop
		default:
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println("\n----------")
}

type item struct {
	price float32
}

func printForRange(items []item) {
	for _, i := range items {
		i.price *= 1.1
	}
	fmt.Printf("%+v\n", items)
	fmt.Println("----------")

	for i := range items {
		items[i].price *= 1.1
	}
	fmt.Printf("%+v\n", items)
	fmt.Println("----------")
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

	// e := 3
	// printForInfBreak(e, 5)

	printForSwitch(10)

	// 構造体3のitem型スライス
	items := []item{
		{price: 10.},
		{price: 20.},
		{price: 30.},
	}

	printForRange(items)
}
