package model

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
	"task05/dao"
)

const (
	PasswordPaper string = "TyroneSalt123!"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null;size:255" json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `gorm:"size:64" json:"email"`
	Posts    []Post
}

func RegisterUser(user *User) (err error) {
	var userDB User
	// 检查用户名是否已存在
	err = dao.DB.Where("username = ?", user.Username).First(&userDB).Error
	if err == nil {
		return errors.New("user already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// 加密密码
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return err
	}
	// 创建新用户
	if err = dao.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func Login(user *User) (err error) {
	var userDB User

	// 查询数据库中是否存在该用户
	err = dao.DB.Where("username = ?", user.Username).First(&userDB).Error
	if err != nil {
		return errors.New("user not found")
	}
	// 加密密码
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return err
	}
	// 检查密码是否匹配
	if userDB.Password != user.Password {
		return errors.New("incorrect password")
	}
	// 密码正确，返回用户信息
	*user = userDB
	return nil
}

// 加密密码
func hashPassword(password string) (string, error) {
	passwordWithPepper := password + PasswordPaper
	hash := sha256.New()
	hash.Write([]byte(passwordWithPepper))
	hashedBytes := hash.Sum(nil)
	hashedPassword := hex.EncodeToString(hashedBytes)
	return hashedPassword, nil
}
