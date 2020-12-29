package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	folderDeleteUID string
)

var folderDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a folder",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.FolderDelete(folderDeleteUID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	folderDeleteCmd.Flags().StringVar(&folderDeleteUID, "uid", "", "Folder UID")
	folderDeleteCmd.MarkFlagRequired("uid")

	folderCmd.AddCommand(folderDeleteCmd)
}
