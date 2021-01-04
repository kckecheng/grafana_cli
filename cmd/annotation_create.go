package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/kckecheng/grafana_cli/client"
	"github.com/spf13/cobra"
)

var (
	annotationCreateDashboardID uint64
	annotationCreatePanelIDs    []uint
	annotationCreateFromTime    string
	annotationCreateToTime      string
	annotationCreateTags        []string
	annotationCreateText        string
)

var annotationCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an annotation",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()

		var from, to time.Time
		var err error
		now := time.Now()
		errfmt := "Cannot parse timestamp %s, please use RFC3339 format like 2020-12-29T15:47:35+08:00\n"
		if annotationCreateFromTime != "" {
			from, err = client.ParseRFC3339TimeString(annotationCreateFromTime)
			if err != nil {
				fmt.Printf(errfmt, annotationCreateFromTime)
				os.Exit(1)
			}
		} else {
			from = now
		}

		if annotationCreateToTime != "" {
			to, err = client.ParseRFC3339TimeString(annotationCreateToTime)
			if err != nil {
				fmt.Printf(errfmt, annotationCreateToTime)
				os.Exit(1)
			}

		} else {
			to = now
		}

		errFlag := false
		var panelIDs []uint64
		if len(annotationCreatePanelIDs) > 0 {
			for _, id := range annotationCreatePanelIDs {
				panelIDs = append(panelIDs, uint64(id))
			}
		} else {
			// if not panel IDs are specified, use all panels of the same dashboard
			dashboards, err := c.DashboardList()
			if err != nil {
				fmt.Println("Fail to query dashbord")
				os.Exit(1)
			}

			var uid string
			for _, dashboard := range dashboards {
				if dashboard.ID == annotationCreateDashboardID {
					uid = dashboard.UID
					break
				}
			}
			if uid == "" {
				fmt.Println("Could not find a dashboard with ID ", annotationCreateDashboardID)
				os.Exit(1)
			}
			panels, err := c.DashboardPanelList(uid)
			if err != nil {
				fmt.Println("Fail to list panels for the specified dashboard")
				os.Exit(1)
			}
			for _, panel := range panels {
				panelIDs = append(panelIDs, panel.ID)
			}

		}
		for _, id := range panelIDs {
			err := c.AnnotationCreate(annotationCreateDashboardID, id, from, to, annotationCreateText, annotationCreateTags...)
			if err != nil {
				errFlag = true
				fmt.Printf("Fail to add annotation to dashboard %v panel %v due to %s", annotationCreateDashboardID, id, err.Error())
			}
		}
		if errFlag {
			os.Exit(1)
		}
	},
}

func init() {
	annotationCreateCmd.Flags().Uint64Var(&annotationCreateDashboardID, "dashboard", 0, "Dashboard ID")
	annotationCreateCmd.Flags().UintSliceVar(&annotationCreatePanelIDs, "panel", []uint{}, "Panel IDs as --panel id1,id2,... or --panel id1 --panel id2 ... (all panels will be used if not specified)")
	annotationCreateCmd.Flags().StringVar(&annotationCreateFromTime, "from", "", "From timestamp (RFC3339, such as 2021-01-10T21:00:00+08:00)")
	annotationCreateCmd.Flags().StringVar(&annotationCreateToTime, "to", "", "To timestamp (RFC3339)")
	annotationCreateCmd.Flags().StringSliceVar(&annotationCreateTags, "tag", []string{}, "Annotation tags as --tag tag1,tag2,... or --tag tag1 --tag tag2 ...")
	annotationCreateCmd.Flags().StringVar(&annotationCreateText, "text", "", "Annotation words")
	annotationCreateCmd.MarkFlagRequired("dashboard")
	annotationCreateCmd.MarkFlagRequired("text")

	annotationCmd.AddCommand(annotationCreateCmd)
}
