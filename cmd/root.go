package cmd

import (
	"fmt"
	"github.com/giffgaff/ipsec_exporter/exporter"
	"github.com/spf13/cobra"
	"os"
)

const (
	flagIpsecConfigFile  = "config-path"
	flagWebListenAddress = "web.listen-address"
)

var Version string
var RootCmd = &cobra.Command{
	Use:   "ipsec_exporter",
	Short: "Prometheus exporter for ipsec status.",
	Long:  "",
	Run:   defaultCommand,
	Version: Version,
}

func init() {
	RootCmd.PersistentFlags().StringVar(&exporter.IpSecConfigFile, flagIpsecConfigFile,
		"/etc/strongswan/ipsec.conf",
		"Path to the ipsec config file.")

	RootCmd.PersistentFlags().IntVar(&exporter.WebListenAddress, flagWebListenAddress,
		9536,
		"Address on which to expose metrics.")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func defaultCommand(_ *cobra.Command, _ []string) {
	exporter.Serve()
}
