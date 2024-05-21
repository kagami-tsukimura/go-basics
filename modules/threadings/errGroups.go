package threadings

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func doTask(ctx context.Context, task string) error {
	var t *time.Ticker
	switch task {
	case "fake1":
		t = time.NewTicker(500 * time.Millisecond)
	case "fake2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(1000 * time.Millisecond)
	}

	select {
	case <-ctx.Done():
		fmt.Printf("%v cancelled: %v\n", task, ctx.Err())
		return ctx.Err()
	case <-t.C:
		t.Stop()
		if task == "fake1" || task == "fake2" {
			return fmt.Errorf("%v process failed", task)
		}
		fmt.Printf("%v done\n", task)
	}
	return nil
}

func ErrGroups() {
	// 初回エラー時にキャンセル
	eg, ctx := errgroup.WithContext(context.Background())
	s := []string{"task1", "fake1", "task2", "fake2", "task3"}

	for _, v := range s {
		task := v
		eg.Go(func() error {
			return doTask(ctx, task)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("all tasks done")
}
