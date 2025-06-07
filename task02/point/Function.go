package main

func addTen(num *int) {
	*num += 10
}

func mulSlice(sp *[]int) {
	for i := range *sp {
		(*sp)[i] *= 2
	}
}
