package cmd

import (
	"github.com/spf13/cobra"
)

var folderCmd = &cobra.Command{
	Use:   "folder",
	Short: "Operate folders",
	Long:  `List, create, and delete folders`,
	// Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(folderCmd)
}
