package main

import (
	"fmt"
	"sync"
	"time"
)

func channelMsgBuff() {
	group := sync.WaitGroup{}
	group.Add(2)
	channel := make(chan int, 4)

	go func() {
		defer group.Done()
		for i := 0; i < 100; i++ {
			channel <- i
		}
		close(channel)
	}()

	go func() {
		defer group.Done()
		timeout := time.After(2 * time.Second)
		for {
			select {
			case msg, ok := <-channel:
				if !ok {
					println("Channel closed, exiting goroutine.")
					return
				}
				println("Received:", msg)
			case <-timeout:
				println("Timeout reached, exiting goroutine.")
				return
			default:
				fmt.Println("没有数据，等待中...")
				time.Sleep(500 * time.Millisecond)
			}

		}
		println("Channel closed, exiting goroutine.")
	}()
	group.Wait()
}
