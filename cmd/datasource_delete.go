package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	datasourceDeleteName string
)

var datasourceDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a data source",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		err := c.DataSourceDelete(datasourceDeleteName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	datasourceDeleteCmd.Flags().StringVar(&datasourceDeleteName, "name", "", "Data source name")
	datasourceDeleteCmd.MarkFlagRequired("name")

	datasourceCmd.AddCommand(datasourceDeleteCmd)
}
