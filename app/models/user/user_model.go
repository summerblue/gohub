// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/logger"
)

// User 用户模型
type User struct {
	models.CommonIDField

	Name     string `json:"name,omitempty"`
	Email    string `gorm:"default:null" json:"-"`
	Phone    string `gorm:"default:null" json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() (err error) {
	if err = database.DB.Create(&userModel).Error; err != nil {
		logger.LogIf(err)
		return err
	}

	return nil
}
