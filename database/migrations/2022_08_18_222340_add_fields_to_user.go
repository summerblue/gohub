package migrations

import (
	"database/sql"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		City         string `gorm:"type:varchar(10);"`
		Introduction string `gorm:"type:varchar(255);"`
		Avatar       string `gorm:"type:varchar(255);default:null"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})

	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&User{}, "City")
		migrator.DropColumn(&User{}, "Introduction")
		migrator.DropColumn(&User{}, "Avatar")
	}

	migrate.Add("2022_08_18_222340_add_fields_to_user", up, down)
}
