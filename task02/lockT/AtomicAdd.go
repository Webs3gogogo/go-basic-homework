package main

import (
	"sync"
	"sync/atomic"
)

func atomicAdd() {
	var (
		group sync.WaitGroup
		sum   int64
	)
	group.Add(10)
	for range 10 {
		go func() {
			defer group.Done()
			for i := range 1000 {
				atomic.AddInt64(&sum, int64(i))
			}
		}()
	}
	group.Wait()
	println(sum)
}
