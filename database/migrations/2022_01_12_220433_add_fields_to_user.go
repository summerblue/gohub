package migrations

import (
	"database/sql"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		City          string `gorm:"type:varchar(10);"`
		Indtroduction string `gorm:"type:varchar(255);"`
		Avatar        string `gorm:"type:varchar(255);default:null"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})

	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&User{}, "City")
		migrator.DropColumn(&User{}, "Indtroduction")
		migrator.DropColumn(&User{}, "Avatar")
	}

	migrate.Add("2022_01_12_220433_add_fields_to_user", up, down)
}
