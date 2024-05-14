/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"path"

	"github.com/spf13/cobra"
	"github.com/zumosik/bb-migrator/pkg"
)

const (
	dirFlag = "dir"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new migration file",
	Long:  `Creates new migration file in the migrations directory.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if path.Ext(name) != ".sql" {
			name = name + ".sql"
		}

		folder, err := cmd.PersistentFlags().GetString(dirFlag)
		if err != nil {
			pkg.L.Printf("cant get %s: %v", dirFlag, err)
			return
		}
		fullPath := path.Join(folder, name)

		pkg.CreateMigrationCtx(context.Background(), pkg.L, fullPath)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().String(dirFlag, "./", "Directory to store migration files")
}
