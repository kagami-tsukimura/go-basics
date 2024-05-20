package threadings

import (
	"sync"
	"time"
)

func Selects() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		ch1 <- "A"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(800 * time.Millisecond)
		ch2 <- "B"
	}()
}
