package cmd

import (
	"os"
)

func init() {
	rootCmd.Flags().StringArrayVarP(
		&filter,
		"filter",
		"f",
		[]string{},
		"Filter Crypto by symbol",
	)

	saveCmd.Flags().StringVarP(
		&fileType,
		"type",
		"t",
		"csv",
		"Determine the type of file that will be created to persist the cryptocurrency data, a CSV or JSON file.",
	)

	saveCmd.Flags().StringVarP(
		&fileName,
		"name",
		"n",
		"criptomine",
		"Define the name of the file containing the cryptocurrency data",
	)

	rootCmd.AddCommand(saveCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
