package main

import "sync"

func channelMsg() {
	group := sync.WaitGroup{}
	group.Add(2)
	channel := make(chan int)

	go func() {
		defer group.Done()
		for i := 0; i < 10; i++ {
			channel <- i
		}
		close(channel)
	}()

	go func() {
		defer group.Done()
		for msg := range channel {
			println("Received:", msg)
		}
		println("Channel closed, exiting goroutine.")
	}()
	group.Wait()
}
