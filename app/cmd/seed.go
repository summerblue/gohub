package cmd

import (
	"gohub/database/seeders"
	"gohub/pkg/console"
	"gohub/pkg/seed"

	"github.com/spf13/cobra"
)

var CmdDBSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the database",
	Run:   runSeeders,
	Args:  cobra.MaximumNArgs(1), // 只允许 1 个参数
}

func runSeeders(cmd *cobra.Command, args []string) {
	seeders.Initialize()
	if len(args) > 0 {
		// 有传参数的情况
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else {
		// 默认运行全部迁移
		seed.RunAll()
		console.Success("Done seeding.")
	}
}
