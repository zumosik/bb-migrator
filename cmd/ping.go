/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/zumosik/bb-migrator/pkg"
)

const (
	driverFlag  = "driver"
	driverPFlag = "D"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Pings database to check connection",
	Long: `Pings database to check connection
	Example: bb-migrator ping -D postgres "postgres://user:password@localhost:5432/dbname"
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uri := args[0]
		driver, err := cmd.PersistentFlags().GetString("driver")
		if err != nil {
			pkg.L.Printf("cant get %s: %v", driverFlag, err)
			return
		}
		if driver == "" {
			pkg.L.Printf("driver (--driver or -D) is required")
			return
		}

		err = pkg.PingDB(context.Background(), driver, uri)
		if err != nil {
			pkg.L.Printf("Failed to ping database: %v", err)
			return
		}
		pkg.L.Printf("Database pinged successfully")
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	pingCmd.PersistentFlags().StringP(driverFlag, driverPFlag, "", "Driver of your database (postgres, mysql, sqlite3)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
