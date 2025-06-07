package main

import "sync"

func addShared() {
	var (
		group = sync.WaitGroup{}
		mutex = sync.Mutex{}
	)
	group.Add(10)
	sum := 0
	for range 10 {
		go func(sum *int) {
			defer group.Done()
			for i2 := range 1000 {
				mutex.Lock()
				*sum += i2
				mutex.Unlock()
			}
		}(&sum)
	}
	group.Wait()

	println(sum)

}
