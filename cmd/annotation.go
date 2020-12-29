package cmd

import (
	"github.com/spf13/cobra"
)

var annotationCmd = &cobra.Command{
	Use:   "annotation",
	Short: "Operate annotations",
	Long:  `List, create, and delete annotations`,
	// Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(annotationCmd)
}
