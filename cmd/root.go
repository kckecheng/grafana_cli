package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kckecheng/grafana_cli/client"
	"github.com/spf13/cobra"
)

var (
	grafanaServer   string
	grafanaUser     string
	grafanaPassword string
)

// rootCmd CLI root command
var rootCmd = &cobra.Command{
	Use:   "grafana_cli",
	Short: "Operate Grafana from CLI",
	Long: `Operate Grafana objects, including data sources, folders,
dashboards, annotations, snapshots, etc., from CLI`,
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Make connection to Grafana
func connectGrafana() client.Client {
	// construct base URL which looks like http[s]://<IP>:<port>/api
	base := strings.TrimRight(grafanaServer, "/")
	if strings.HasPrefix(base, "http://") || strings.HasPrefix(base, "https://") {
		base = base + "/api"
	} else {
		base = fmt.Sprintf("http://%s/api", base)
	}

	// init Client
	c := client.New(base, grafanaUser, grafanaPassword)

	// Make connection to Grafana
	err := c.Login()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return c
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&grafanaServer, "server", "http://localhost:3000", "Grafana server (GRAFANA_SERVER)")
	rootCmd.PersistentFlags().StringVar(&grafanaUser, "user", "admin", "Grafana user name (GRAFANA_USER)")
	rootCmd.PersistentFlags().StringVar(&grafanaPassword, "password", "admin", "Grafana user password (GRAFANA_PASSWORD)")

	// Overwrite Grafana server/user/password if env vars are defined
	s, existed := os.LookupEnv("GRAFANA_SERVER")
	if existed {
		grafanaServer = s
	}
	u, existed := os.LookupEnv("GRAFANA_USER")
	if existed {
		grafanaUser = u
	}
	p, existed := os.LookupEnv("GRAFANA_PASSWORD")
	if existed {
		grafanaPassword = p
	}
}
