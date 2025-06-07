package main

import (
	"fmt"
	"sync"
	"time"
)

func addTaskCaller(functions []func()) map[string]int {
	var group sync.WaitGroup
	group.Add(len(functions))
	taskDuration := make(map[string]int)

	for i := range functions {
		go func(index int, f func()) {
			defer group.Done()
			start := time.Now()
			f() // Example arguments, can be modified as needed
			duration := time.Since(start)
			taskDuration[fmt.Sprintf("TaskDuration-%d", i)] = int(duration.Milliseconds())
		}(i, functions[i])
	}
	group.Wait()
	return taskDuration
}
