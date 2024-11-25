package cmd

import (
	"crypto-mine-cli/cmd/commands/save"
	"crypto-mine-cli/cmd/commands/scrape"

	"github.com/spf13/cobra"
)

var (
	saveResults               bool
	fileType                  string
	symbolFilter, compareList []string
)

var (
	rootCmd = &cobra.Command{
		Use:   "crypto-mine",
		Short: "crypto-mine-cli collects cryptocurrency data on the coin market cape website",
		Long: `Crypto crypto-mine-cli is an application that extracts the
data of various cryptocurrencies from the Coin Market Cap
		`,
		Run: func(cmd *cobra.Command, args []string) {
			scrape.Scrape(symbolFilter)
		},
	}

	saveCmd = &cobra.Command{
		Use:   "save",
		Short: "Save stores the results in Donwloads folder",
		Long:  "Save stores the results in Downloads folder, The information from the table is applied to the file of your choice, choose between JSON or CSV with the --type flag (or -t for short)",
		Run: func(cmd *cobra.Command, args []string) {
			save.Save(fileType)
		},
	}
)