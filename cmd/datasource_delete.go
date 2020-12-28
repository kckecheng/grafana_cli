package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dsdname string // data source delete name
)

var datasourceDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a data source",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DataSourceDelete(dsdname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	datasourceDeleteCmd.Flags().StringVar(&dsdname, "name", "", "Data source name")
	datasourceDeleteCmd.MarkFlagRequired("name")

	datasourceCmd.AddCommand(datasourceDeleteCmd)
}
