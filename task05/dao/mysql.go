package dao

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"task05/setting"
)

var (
	DB *gorm.DB // 全局数据库连接对象
)

const DsnTemplate = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

// InitDataBase 初始化数据库连接
func InitDataBase(config *setting.DatabaseConfig) error {
	// 拼接 DSN
	dsn := fmt.Sprintf(DsnTemplate, config.Username, config.Password, config.Host, config.Port, config.Database)

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// 获取底层的 *sql.DB 对象
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Failed to get underlying sql.DB: %v", err)
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 设置最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接的最大生命周期

	// 赋值全局变量
	DB = db
	log.Println("Database connection established successfully")
	return nil
}

// Close 关闭数据库连接
func Close() error {
	if DB == nil {
		return nil // 如果数据库未初始化，则直接返回
	}

	// 获取底层的 *sql.DB 对象
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Failed to get underlying sql.DB: %v", err)
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// 关闭数据库连接
	err = sqlDB.Close()
	if err != nil {
		log.Printf("Failed to close database connection: %v", err)
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	log.Println("Database connection closed successfully")
	return nil
}
