package threadings

import (
	"context"
	"fmt"
	"time"
)

func Contexts() {
	// Context: メインgoroutineからサブgoroutineを一括キャンセル
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	fmt.Println(ctx.Err())
}
