package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	deuid   string // dashboard export UID
	defpath string // dashboard export file path
)

// dashboardCmd represents the dashboard command
var dashboardExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DashboardExport(defpath, deuid)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	dashboardExportCmd.Flags().StringVar(&deuid, "uid", "", "Dashboard UID")
	dashboardExportCmd.Flags().StringVar(&defpath, "path", "", "Dashboard exported file path")
	dashboardExportCmd.MarkFlagRequired("uid")
	dashboardExportCmd.MarkFlagRequired("path")

	dashboardCmd.AddCommand(dashboardExportCmd)
}
