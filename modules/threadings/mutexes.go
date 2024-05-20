package threadings

import (
	"fmt"
	"sync"
	"time"
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

func rwMutexes() {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	var c int
	read(&mu, &wg, &c)
}

func read(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("read lock")
	fmt.Println("c: ", *c)
	time.Sleep(10 * time.Second)
	fmt.Println("read unlock")
}

func Mutexes() {
	raceMutexes()
	rwMutexes()
}
