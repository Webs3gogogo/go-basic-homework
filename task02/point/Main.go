package main

import "fmt"

func main() {
	i := 5
	fmt.Println(i)
	addTen(&i)
	fmt.Println(i)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	mulSlice(&slice)
	fmt.Println(slice)
}
