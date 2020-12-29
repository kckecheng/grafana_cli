package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var snapshotListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snapshots",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		snapshots, err := c.SnapshotList()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%-35v %s\n", "Key", "Name")
		for _, snapshot := range snapshots {
			fmt.Printf("%-35v %s\n", snapshot.Key, snapshot.Name)
		}
	},
}

func init() {
	snapshotCmd.AddCommand(snapshotListCmd)
}
