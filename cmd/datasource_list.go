package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var datasourceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all data sources",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()
		datasources, err := c.DataSourceList()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%-20v %-30v %s\n", "Type", "Name", "URL")
		for _, datasource := range datasources {
			fmt.Printf("%-20v %-30v %s\n", datasource.Type, datasource.Name, datasource.URL)
		}
	},
}

func init() {
	datasourceCmd.AddCommand(datasourceListCmd)
}
