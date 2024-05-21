package threadings

import (
	"context"
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
	heartBeat, out := tasks(ctx, beatInterval)
}
