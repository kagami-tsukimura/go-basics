package threadings

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func Tracers() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error: ", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalln("Error: ", err)
	}
	defer trace.Stop()

	ctx, t := trace.NewTask(context.Background(), "task")
	defer t.End()
	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

	// 逐次処理
	task(ctx, "Task1")
	task(ctx, "Task2")
	task(ctx, "Task3")

}

func task(ctx context.Context, name string) {
	defer trace.StartRegion(ctx, name).End()
	time.Sleep(time.Second)
	fmt.Println(name)
}

func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}
