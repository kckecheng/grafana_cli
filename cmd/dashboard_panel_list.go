package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dashboardPanelListUID string
)

var dashboardPanelListCmd = &cobra.Command{
	Use:   "panel",
	Short: "List panels of a dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		panels, err := c.DashboardPanelList(dashboardPanelListUID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%-5v %s\n", "ID", "Title")
		for _, panel := range panels {
			fmt.Printf("%-5v %s\n", panel.ID, panel.Title)
		}
	},
}

func init() {
	dashboardPanelListCmd.Flags().StringVar(&dashboardPanelListUID, "uid", "", "Dashboard UID")
	dashboardPanelListCmd.MarkFlagRequired("uid")

	dashboardCmd.AddCommand(dashboardPanelListCmd)
}
