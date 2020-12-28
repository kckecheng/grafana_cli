package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dsiname string // data source import name
	dsipath string // data source import file path
)

var datasourceImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a data source",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DataSourceImport(dsiname, dsipath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	datasourceImportCmd.Flags().StringVar(&dsiname, "name", "", "Data source name")
	datasourceImportCmd.Flags().StringVar(&dsipath, "path", "", "Data source exported file path")
	datasourceImportCmd.MarkFlagRequired("name")
	datasourceImportCmd.MarkFlagRequired("path")

	datasourceCmd.AddCommand(datasourceImportCmd)
}
