package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dashboardExportUID  string
	dashboardExportPath string
)

var dashboardExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DashboardExport(dashboardExportPath, dashboardExportUID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	dashboardExportCmd.Flags().StringVar(&dashboardExportUID, "uid", "", "Dashboard UID")
	dashboardExportCmd.Flags().StringVar(&dashboardExportPath, "path", "", "Dashboard exported file path")
	dashboardExportCmd.MarkFlagRequired("uid")
	dashboardExportCmd.MarkFlagRequired("path")

	dashboardCmd.AddCommand(dashboardExportCmd)
}
