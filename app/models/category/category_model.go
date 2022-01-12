//Package category 模型
package category

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Category struct {
	models.BaseModel

	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	models.CommonTimestampsField
}

func (category *Category) Create() {
	database.DB.Create(&category)
}

func (category *Category) Save() (rowsAffected int64) {
	result := database.DB.Save(&category)
	return result.RowsAffected
}

func (category *Category) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&category)
	return result.RowsAffected
}
