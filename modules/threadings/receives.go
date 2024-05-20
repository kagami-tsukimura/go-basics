package threadings

import (
	"fmt"
	"sync"
	"time"
)

func Packages() {
	fmt.Println("Packages")
}

func countProductor(wg *sync.WaitGroup, ch chan<- int, size int, sleep int) {
	defer wg.Done()
	defer close(ch)

	for i := 0; i < size; i++ {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ch <- i
	}
}
