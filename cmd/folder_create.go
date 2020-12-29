package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	folderCreateName string
)

var folderCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a folder",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.FolderCreate(folderCreateName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	folderCreateCmd.Flags().StringVar(&folderCreateName, "name", "", "Folder name")
	folderCreateCmd.MarkFlagRequired("name")

	folderCmd.AddCommand(folderCreateCmd)
}
