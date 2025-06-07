package main

import "fmt"

func main() {
	circle := &Circle{
		radius: 5.0,
	}

	rectangle := &Rectangle{
		length: 10.0,
		width:  5.0,
	}
	fmt.Println(circle.Area())
	fmt.Println(circle.Perimeter())

	fmt.Println(rectangle.Area())
	fmt.Println(rectangle.Perimeter())

	e := &Employee{
		Person: Person{
			Name: "John Doe",
			Age:  "30",
		},
		EmployeeID: 1,
	}
	e.PrintInfo()
}
