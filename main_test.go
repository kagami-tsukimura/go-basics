package main

import (
	"go-basics/modules/threadings"
	"testing"

	"go.uber.org/goleak"
)

func TestLeak(t *testing.T) {
	// goroutine leakのテスト
	defer goleak.VerifyNone(t)
	threadings.Channels()
}
