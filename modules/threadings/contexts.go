package threadings

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func taskTimeout(wg *sync.WaitGroup) {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	wg.Add(3)
	go subTask(ctx, wg, 1)
	go subTask(ctx, wg, 2)
	go subTask(ctx, wg, 3)
	wg.Wait()
}

func subTask(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	t := time.NewTicker(500 * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	case <-t.C:
		t.Stop()
		fmt.Println(id)
	}
}

func taskCancel(wg *sync.WaitGroup) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(1)

	// goroutine criticalTask
	go func() {
		defer wg.Done()
		v, err := criticalTask(ctx)
		if err != nil {
			fmt.Printf("critical task cancelled due to: %v\n", err)
			cancel()
			return
		}
		fmt.Println("success: ", v)
	}()

	wg.Add(1)
	// goroutine normalTask
	go func() {
		defer wg.Done()
		v, err := normalTask(ctx)
		if err != nil {
			fmt.Printf("normal task cancelled due to: %v\n", err)
			return
		}
		fmt.Println("success: ", v)
	}()

	wg.Wait()
}

func normalTask(ctx context.Context) (string, error) {
	t := time.NewTicker(3000 * time.Millisecond)

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
	}
	return "OK_Normal", nil
}

func criticalTask(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()
	t := time.NewTicker(1000 * time.Millisecond)

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
	}
	return "OK_Critical", nil
}

func Contexts() {
	// Context: メインgoroutineからサブgoroutineを一括キャンセル
	var wg sync.WaitGroup
	taskTimeout(&wg)
	taskCancel(&wg)

}
