package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string
	Department string
	Salary     float64
}

type Book struct {
	gorm.Model
	Title  string
	Author string
	Price  float64
}

type CustomEmployee struct {
	Name       string
	Department string
}
