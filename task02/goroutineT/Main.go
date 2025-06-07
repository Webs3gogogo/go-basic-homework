package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//var group sync.WaitGroup
	//group.Add(2)
	//printfForGoroutine(&group)
	//group.Wait()
	//fmt.Println("execution completed")

	task := []func(){
		func() {
			rand.Seed(time.Now().UnixNano())
			min := 100
			max := 500
			delay := rand.Intn(max-min+1) + min
			time.Sleep(time.Duration(delay) * time.Millisecond)
		},
		func() {
			rand.Seed(time.Now().UnixNano())
			min := 100
			max := 500
			delay := rand.Intn(max-min+1) + min
			time.Sleep(time.Duration(delay) * time.Millisecond)
		},
		func() {
			rand.Seed(time.Now().UnixNano())
			min := 100
			max := 500
			delay := rand.Intn(max-min+1) + min
			time.Sleep(time.Duration(delay) * time.Millisecond)
		},
		func() {
			rand.Seed(time.Now().UnixNano())
			min := 100
			max := 500
			delay := rand.Intn(max-min+1) + min
			time.Sleep(time.Duration(delay) * time.Millisecond)
		},
	}

	result := addTaskCaller(task)
	fmt.Println(result)

}
