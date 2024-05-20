package threadings

import (
	"fmt"
	"sync"
)

func Mutexes() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var i int
	wg.Add(2)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		i++
	}()
	go func() {
		defer wg.Done()
		i++
	}()
	wg.Wait()
	fmt.Println(i)
}
