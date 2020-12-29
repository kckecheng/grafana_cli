package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	datasourceImportName string
	datasourceImportPath string
)

var datasourceImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a data source",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DataSourceImport(datasourceImportName, datasourceImportPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	datasourceImportCmd.Flags().StringVar(&datasourceImportName, "name", "", "Data source name")
	datasourceImportCmd.Flags().StringVar(&datasourceImportPath, "path", "", "Data source exported file path")
	datasourceImportCmd.MarkFlagRequired("name")
	datasourceImportCmd.MarkFlagRequired("path")

	datasourceCmd.AddCommand(datasourceImportCmd)
}
