package user

import (
	"gohub/pkg/hash"

	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
