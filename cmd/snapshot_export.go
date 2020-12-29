package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	snapshotExportKey  string
	snapshotExportPath string
)

var snapshotExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a snapshot",
	Long:  "Export a snapshot which can be imported as a dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.SnapshotExport(snapshotExportKey, snapshotExportPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	snapshotExportCmd.Flags().StringVar(&snapshotExportKey, "key", "", "Snapshot key")
	snapshotExportCmd.Flags().StringVar(&snapshotExportPath, "path", "", "Snapshot export path")
	snapshotExportCmd.MarkFlagRequired("key")
	snapshotExportCmd.MarkFlagRequired("path")

	snapshotCmd.AddCommand(snapshotExportCmd)
}
