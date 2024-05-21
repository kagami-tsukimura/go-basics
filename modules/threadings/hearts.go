package threadings

import (
	"context"
	"io"
	"log"
	"os"
	"time"
)

func task(ctx context.Context, beatInterval time.Duration) (<-chan struct{}, <-chan time.Time) {

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
}
