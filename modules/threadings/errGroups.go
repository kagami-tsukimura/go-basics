package threadings

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func doTask(task string) error {
	if task == "fake1" || task == "fake2" {
		return fmt.Errorf("%v failed", task)
	}
	fmt.Printf("task %v done\n", task)
	return nil
}

func ErrGroups() {
	eg := new(errgroup.Group)
	s := []string{"task1", "fake1", "task2", "fake2"}

	for _, v := range s {
		task := v
		eg.Go(func() error {
			return doTask(task)
		})
	}
}
