package cmd

import (
	"github.com/spf13/cobra"
)

var datasourceCmd = &cobra.Command{
	Use:   "datasource",
	Short: "operate data sources",
	Long:  `List, export, import and delete data sources`,
	// Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(datasourceCmd)
}
