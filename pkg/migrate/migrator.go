// Package migrate 处理数据库迁移
package migrate

import (
	"gohub/pkg/database"

	"gorm.io/gorm"
)

// Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 对应数据的 migrations 表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// NewMigrator 创建 Migrator 实例，用以执行迁移操作
func NewMigrator() *Migrator {

	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	// migrations 不存在的话就创建它
	migrator.createMigrationsTable()

	return migrator
}

// 创建 migrations 表
func (migrator *Migrator) createMigrationsTable() {

	migration := Migration{}

	// 不存在才创建
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}
