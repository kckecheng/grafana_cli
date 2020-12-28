package cmd

import (
	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "operate dashboards",
	Long:  `List, export, import and delete dashboards`,
	// Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
