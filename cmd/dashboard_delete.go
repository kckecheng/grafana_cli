package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dduid string // dashboard delete uid
)

// dashboardCmd represents the dashboard command
var dashboardDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DashboardDelete(dduid)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	dashboardDeleteCmd.Flags().StringVar(&dduid, "uid", "", "Dashboard UID for deletion")
	dashboardDeleteCmd.MarkFlagRequired("uid")

	dashboardCmd.AddCommand(dashboardDeleteCmd)
}
