package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	snapshotDeleteKey string
)

var snapshotDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.SnapshotDelete(snapshotDeleteKey)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	snapshotDeleteCmd.Flags().StringVar(&snapshotDeleteKey, "key", "", "Snapshot key")
	snapshotDeleteCmd.MarkFlagRequired("key")

	snapshotCmd.AddCommand(snapshotDeleteCmd)
}
