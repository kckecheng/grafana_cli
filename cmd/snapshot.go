package cmd

import (
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Operate snapshots",
	Long:  `List, export, and delete snapshots`,
	// Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(snapshotCmd)
}
