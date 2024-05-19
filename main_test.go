package main

import (
	"go-basics/modules/threadings"
	"testing"

	"go.uber.org/goleak"
)

func TestLeak(t *testing.T) {
	// goroutine leakのテスト
	defer goleak.VerifyNone(t)
	// goroutine leakしている関数を直接呼び出す
	threadings.Channels()
}
