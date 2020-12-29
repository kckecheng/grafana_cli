package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dashboardImportPath     string
	dashboardImportFolderID uint64
)

var dashboardImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DashboardImport(dashboardImportPath, dashboardImportFolderID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	dashboardImportCmd.Flags().StringVar(&dashboardImportPath, "path", "", "Dashboard file path")
	dashboardImportCmd.Flags().Uint64Var(&dashboardImportFolderID, "folder", 0, `Folder ID for the imported dashboard, "0" for general`)
	dashboardImportCmd.MarkFlagRequired("folder")
	dashboardImportCmd.MarkFlagRequired("path")

	dashboardCmd.AddCommand(dashboardImportCmd)
}
