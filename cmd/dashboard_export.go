package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	deuid  string // dashboard export UID
	depath string // dashboard export file path
)

var dashboardExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DashboardExport(depath, deuid)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	dashboardExportCmd.Flags().StringVar(&deuid, "uid", "", "Dashboard UID")
	dashboardExportCmd.Flags().StringVar(&depath, "path", "", "Dashboard exported file path")
	dashboardExportCmd.MarkFlagRequired("uid")
	dashboardExportCmd.MarkFlagRequired("path")

	dashboardCmd.AddCommand(dashboardExportCmd)
}
