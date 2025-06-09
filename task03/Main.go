package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
	Type  uint `gorm:"default:1"` // 默认值为 1
}
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	gorm.Model
	Name       string
	CompanyID  uint
	Company    Company
	CreditCard CreditCard
}
type Company struct {
	ID   int
	Name string
}

func GetDbInstance() (*gorm.DB, error) {
	username := "root"
	password := "admin"
	dbName := "BlogForGo"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {

	db, err := GetDbInstance()
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{})
	db.AutoMigrate(&Company{})
	//db.Create(&User{
	//	Name:       "jinzhu",
	//	CreditCard: CreditCard{Number: "411111111111"},
	//})
	db.Create(&User{
		Name:       "jinzhu",
		CreditCard: CreditCard{Number: "411111111111"},
		Company: Company{
			Name: "TestCompany",
		},
	})
	// Create
	//db.Debug().Create(&newProduct)
	//db.Debug().Create(products)
	//db.Debug().Create(&Product{Code: "ACE", Price: 100})
	//db.Debug().Create(&Product{Code: "D42", Price: 200}) // 使用指针创建
	//db.Debug().Omit("Price").Create(&Product{Code: "A42", Price: 100}) // 使用指针创建
	//db.Unscoped().Find(&products)
	//for _, p := range products {
	//	fmt.Println(p)
	//	//fmt.Printf("ID: %d, Code: %s, Price: %d\n", p.ID, p.Code, p.Price)
	//}
	//db.Unscoped().Debug().Where("1 =1 ").Delete(&Product{})
	//fmt.Println(result)
	// SELECT * FROM users WHERE age = 20;

	//db.Debug().Model(&Product{}).Where("code", "A42").UpdateColumn("Price", 200) // 仅更新 Price 字段
	//db.Debug().Delete(&Product{}, "1=1") // 删除 code 字段值为 A42 的记录
	//// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//
	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
