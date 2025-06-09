package main

import "gorm.io/gorm"

type Student struct {
	*gorm.Model
	Name  string
	Age   uint
	Grade string
}

type Account struct {
	*gorm.Model
	Balance float64
}

type Transaction struct {
	*gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
}
