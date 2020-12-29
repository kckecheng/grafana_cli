package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var folderListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all folders",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		folders, err := c.FolderList()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%-5v %-15v %s\n", "ID", "UID", "Name")
		for _, folder := range folders {
			fmt.Printf("%-5v %-15v %s\n", folder.ID, folder.UID, folder.Title)
		}
	},
}

func init() {
	folderCmd.AddCommand(folderListCmd)
}
