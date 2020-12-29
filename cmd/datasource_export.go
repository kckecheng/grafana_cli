package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	datasourceExportName string
	datasourceExportPath string
)

var datasourceExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a data source",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DataSourceExport(datasourceExportName, datasourceExportPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	datasourceExportCmd.Flags().StringVar(&datasourceExportName, "name", "", "Data source name")
	datasourceExportCmd.Flags().StringVar(&datasourceExportPath, "path", "", "Data source exported file path")
	datasourceExportCmd.MarkFlagRequired("name")
	datasourceExportCmd.MarkFlagRequired("path")

	datasourceCmd.AddCommand(datasourceExportCmd)
}
