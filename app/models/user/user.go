package user

import (
	"goblog_2/app/models"
	"goblog_2/pkg/password"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	// gorm:"-" —— 设置 GORM 在读写时略过此字段，仅用于表单验证
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

// ComparePassword 对比密码是否匹配
func (user *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, user.Password)
}
