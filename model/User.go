package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"log"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 检查用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Where("username=?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPW(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 加密密码
func ScryptPW(password string) string {
	const KeyLen = 10
	//salt := make([]byte, 8)
	salt := []byte{12, 21, 21, 22, 23, 11, 33, 65}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

// 编辑用户
func EditUser(id int, data *User) int {

	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&User{}).Where("id=?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 检查用户登录
func CheckLogin(username, password string) int {
	// code:=CheckUser(username)
	// if code==errmsg.ERROR_USER_NOT_EXIST
	var user User
	var code = errmsg.SUCCSE
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		code = errmsg.ERROR_USER_NOT_EXIST
		return code
	}
	if ScryptPW(password) != user.Password {
		code = errmsg.ERROR_PASSWORD_WRONG
		return code
	}
	if user.Role != 0 {
		code = errmsg.ERROR_USER_NO_RIGHT
		return code
	}
	return code
}
