package threadings

import (
	"fmt"
	"sync"
	"sync/atomic"
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
	var rwMu sync.RWMutex
	var c int
	wg.Add(5)
	go write(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	go write(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)

	wg.Wait()
	fmt.Println("finished")
	fmt.Println("----------")
}

func read(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("read lock")
	fmt.Println("c: ", *c)
	time.Sleep(1 * time.Second)
	fmt.Println("read unlock")
}

func write(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("write lock")
	*c += 1
	time.Sleep(1 * time.Second)
	fmt.Println("write unlock")
}

func atomics() {
	var wg sync.WaitGroup
	var c int64

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&c, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(c)
	fmt.Println("----------")
}

func Mutexes() {
	raceMutexes()
	rwMutexes()
	atomics()
}
