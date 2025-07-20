package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDbInstance() (*gorm.DB, error) {
	username := "root"
	password := "admin"
	dbName := "BlogForGo"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	//handelStudentTable()
	handleTransaction()
}

func handelStudentTable() {
	db, _ := GetDbInstance()
	err := db.AutoMigrate(&Student{})

	if err != nil {
		return
	}
	//db.Debug().Create(&Student{
	//	Name:  "张三",
	//	Age:   20,
	//	Grade: "三年级",
	//})

	var students []Student
	db.Debug().Where("age > 18").Find(&students)
	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Grade: %s\n", student.ID, student.Name, student.Age, student.Grade)
	}

	db.Debug().Where("name = ?", "张三").UpdateColumn("grade", "四年级")
	var student Student
	db.Debug().Where("name = ?", "张三").Find(&student)
	fmt.Println(student)

	db.Debug().Where("age < ?", 15).Delete(&Student{})

	db.Debug().Find(&students)
	fmt.Println(students)
}

func handleTransaction() {

	db, _ := GetDbInstance()

	err := db.AutoMigrate(&Account{}, &Transaction{})

	if err != nil {
		return
	}
	accounts := []*Account{
		{
			Balance: 1000.0,
		},
		{
			Balance: 2000.0,
		},
	}

	db.Debug().Create(accounts)

	db.Transaction(func(transferTx *gorm.DB) error {
		var A, B Account
		transferTx.Debug().Where("id = ?", 1).Find(&A)
		if A.Balance < 1000 {
			return errors.New("Insufficient balance in account.png A")
		} else {
			A.Balance -= 1000
		}
		transferTx.Debug().Save(&A)
		transferTx.Debug().Where("id = ?", 2).Find(&B)
		B.Balance += 1000
		transferTx.Debug().Save(&B)

		transcation := Transaction{
			FromAccountID: A.ID,
			ToAccountID:   B.ID,
			Amount:        100,
		}
		transferTx.Debug().Create(&transcation)

		fmt.Printf("Transfer successful: From Account ID %d to Account ID %d, Amount: %.2f\n", A.ID, B.ID, transcation.Amount)

		return nil
	})

}
