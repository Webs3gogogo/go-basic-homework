package main

import (
	"fmt"
	"sync"
)

func printfForGoroutine(group *sync.WaitGroup) {
	go func(threadName string) {
		defer group.Done()
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Printf("%s : %d \n", threadName, i)
			}
		}
	}("goroutineT-even")

	go func(threadName string) {
		defer group.Done()
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				fmt.Printf("%s : %d \n", threadName, i)
			}
		}
	}("goroutineT-odd")

}
