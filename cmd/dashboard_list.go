package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var dashboardListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all dashboards",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		dashboards, err := c.DashboardList()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%-15v %-15v %-15v %s\n", "Folder ID", "Dashboard ID", "Dashboard UID", "Dashboard Title")
		for _, dashboard := range dashboards {
			fmt.Printf("%-15v %-15v %-15v %s\n", dashboard.FolderID, dashboard.ID, dashboard.UID, dashboard.Title)
		}
	},
}

func init() {
	dashboardCmd.AddCommand(dashboardListCmd)
}
