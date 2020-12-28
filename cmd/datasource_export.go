package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dsename string // data source export name
	dsepath string // data source export file path
)

var datasourceExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a data source",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DataSourceExport(dsename, dsepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	datasourceExportCmd.Flags().StringVar(&dsename, "name", "", "Data source name")
	datasourceExportCmd.Flags().StringVar(&dsepath, "path", "", "Data source exported file path")
	datasourceExportCmd.MarkFlagRequired("name")
	datasourceExportCmd.MarkFlagRequired("path")

	datasourceCmd.AddCommand(datasourceExportCmd)
}
