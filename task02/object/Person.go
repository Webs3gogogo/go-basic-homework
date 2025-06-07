package main

import "fmt"

type Person struct {
	Name, Age string
}

type Employee struct {
	Person
	EmployeeID float64
}

func (e *Employee) PrintInfo() {
	fmt.Println("{Name: ", e.Name, ", Age: ", e.Age, ", EmployeeID: ", e.EmployeeID, "}")
}
