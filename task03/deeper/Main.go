package main

import (
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
	handleQuery()
}
func handleQuery() {
	db, _ := GetDbInstance()
	err := db.AutoMigrate(&Employee{})
	if err != nil {
		return
	}

	employees := []*Employee{
		{
			Name:       "Alice Smith",
			Department: "人事部",
			Salary:     60000.00,
		},
		{
			Name:       "Bob Johnson",
			Department: "财务部",
			Salary:     80000.00,
		},
		{
			Name:       "John Doe",
			Department: "技术部",
			Salary:     75000.00,
		},
	}

	db.Debug().Create(&employees)
	var customEmp []CustomEmployee
	db.Debug().Where("department = ?", "技术部").Find(&employees).Scan(&customEmp)

	for _, emp := range customEmp {
		fmt.Printf("Name: %s, Department: %s\n", emp.Name, emp.Department)
	}

}
