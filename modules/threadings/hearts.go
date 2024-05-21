package threadings

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func tasks(ctx context.Context, beatInterval time.Duration) (<-chan struct{}, <-chan time.Time) {
	heartBeat := make(chan struct{})
	out := make(chan time.Time)
	go func() {
		defer close(heartBeat)
		defer close(out)
		pulse := time.NewTicker(beatInterval)
		task := time.NewTicker(2 * beatInterval)

		sendPulse := func() {
			select {
			case heartBeat <- struct{}{}:
			default:
			}
		}

		sendValue := func(t time.Time) {
			for {
				select {
				case <-ctx.Done():
					return
				case <-pulse.C:
					sendPulse()
				case out <- t:
					return
				}
			}
		}

		for {
			select {
			case <-ctx.Done():
				return
			case <-pulse.C:
				sendPulse()
			case t := <-task.C:
				sendValue(t)
			}
		}
	}()
	return heartBeat, out
}

func Hearts() {
	file, err := os.Create("heart.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.LstdFlags)
	ctx, cancel := context.WithTimeout(context.Background(), 5100*time.Millisecond)
	defer cancel()
	const wdtTimeout = 800 * time.Millisecond
	const beatInterval = 500 * time.Millisecond
	heartBeat, v := tasks(ctx, beatInterval)

loop:
	for {
		select {
		case _, ok := <-heartBeat:
			if !ok {
				break loop
			}
			fmt.Println("beat pulse")
		case r, ok := <-v:
			if !ok {
				break loop
			}
			t := strings.Sprit(r.String(), "m=")
			fmt.Printf("value: %v [s]\n", t[1])
		case <-time.After(wdtTimeout):
			errorLogger.Println("wdt timeout")
			break loop
		}
	}
}
