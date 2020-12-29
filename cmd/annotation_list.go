package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/kckecheng/grafana_cli/client"
	"github.com/spf13/cobra"
)

var (
	annotationListDashboardID uint64
	annotationListFromTime    string
	annotationListToTime      string
)

var annotationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List annotations",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()

		var from, to time.Time
		var err error
		errfmt := "Cannot parse timestamp %s, please use RFC3339 format like 2020-12-29T15:47:35+08:00\n"
		if annotationListFromTime != "" {
			from, err = client.ParseRFC3339TimeString(annotationListFromTime)
			if err != nil {
				fmt.Printf(errfmt, annotationListFromTime)
				os.Exit(1)
			}
		} else {
			from = time.Unix(0, 0)
		}

		if annotationListToTime != "" {
			to, err = client.ParseRFC3339TimeString(annotationListToTime)
			if err != nil {
				fmt.Printf(errfmt, annotationListToTime)
				os.Exit(1)
			}

		} else {
			to = time.Now()
		}

		annotations, err := c.AnnotationList(annotationListDashboardID, from, to)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%-15v %-15v %-15v %s\n", "ID", "Dashboard ID", "Panel ID", "Text")
		for _, annotation := range annotations {
			fmt.Printf("%-15v %-15v %-15v %s\n", annotation.ID, annotation.DashboardID, annotation.PanelID, annotation.Text)
		}
	},
}

func init() {
	annotationListCmd.Flags().Uint64Var(&annotationListDashboardID, "dashboard", 0, "Dashboard ID")
	annotationListCmd.Flags().StringVar(&annotationListFromTime, "from", "", "From timestamp (RFC3339)")
	annotationListCmd.Flags().StringVar(&annotationListToTime, "to", "", "To timestamp (RFC3339)")

	annotationCmd.AddCommand(annotationListCmd)
}
