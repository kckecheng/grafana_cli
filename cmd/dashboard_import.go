package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	difpath    string // dashboard import file path
	difolderid uint64 // dashboard import folder ID
)

// dashboardCmd represents the dashboard command
var dashboardImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DashboardImport(difpath, difolderid)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	dashboardImportCmd.Flags().StringVar(&difpath, "path", "", "Dashboard file path")
	dashboardImportCmd.Flags().Uint64Var(&difolderid, "folderid", 0, "Folder ID for the imported dashboard")
	dashboardImportCmd.MarkFlagRequired("folderid")
	dashboardImportCmd.MarkFlagRequired("path")

	dashboardCmd.AddCommand(dashboardImportCmd)
}
