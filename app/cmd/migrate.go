package cmd

import (
	"gohub/database/migrations"
	"gohub/pkg/migrate"

	"github.com/spf13/cobra"
)

var migrator *migrate.Migrator

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	// 所有 migrate 下的子命令都会执行以下代码
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
	)
}

func initMigrator() {
	// 初始化 migrator
	migrator = migrate.NewMigrator()
	// 注册 database/migrations 下的所有迁移文件
	migrations.Initialize()
}

func runUp(cmd *cobra.Command, args []string) {
	initMigrator()
	migrator.Up()
}
