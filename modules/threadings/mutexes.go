package threadings

import (
	"fmt"
	"sync"
)

func raceMutexes() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var i int
	wg.Add(2)
	go func() {
		defer wg.Done()
		// MutexによるLock
		mu.Lock()
		// 処理完了後にUnLock
		defer mu.Unlock()
		// i++
		i = 1
	}()
	go func() {
		defer wg.Done()
		// MutexによるLock
		mu.Lock()
		// 処理完了後にUnLock
		defer mu.Unlock()
		// i++
		i = 2
	}()
	wg.Wait()
	fmt.Println(i)
	fmt.Println("----------")
}

func Mutexes() {
	raceMutexes()
}
