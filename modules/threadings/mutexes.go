package threadings

import (
	"fmt"
	"sync"
)

func Mutexes() {
	var wg sync.WaitGroup
	var i int
	wg.Add(2)
	go func() {
		defer wg.Done()
		i++
	}()
	go func() {
		defer wg.Done()
		i++
	}()
	wg.Wait()
	fmt.Println(i)
}
