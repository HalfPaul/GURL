package cmd

import (
	"fmt"

	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"GURL/protocols"
)

type Result struct{}

var headerVar string
var protocolVar string
var dataVar string

var rootCmd = &cobra.Command{
	Use:   "GURL",
	Short: "CURL alternative",
	Long:  `Command-line tool for transferring data using various network protocols.`,
	Run: func(cmd *cobra.Command, args []string) {

		switch protocolVar {
		case "":
			color.Green(protocols.GetData(args[0], headerVar))
		case "POST":
			if dataVar == "" {
				color.Red("No data was given for POST request.")
			}
			color.Green(protocols.PostData(args[0], dataVar, headerVar))
		case "PUT":
			if dataVar == "" {
				color.Red("No data was given for PUT request.")
			}
			color.Green(protocols.PutData(args[0], dataVar, headerVar))
		case "GET":
			color.Green(protocols.GetData(args[0], headerVar))
		default:
			color.Red(fmt.Sprintf("Protocol by the name %s doesn't exist", protocolVar))
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringVarP(&headerVar, "header", "H", "", "Variable for passing header.")
	rootCmd.Flags().StringVarP(&protocolVar, "protocol", "X", "", "Variable to choose protocol.")
	rootCmd.Flags().StringVarP(&dataVar, "data", "d", "", "Variable for data.")
}
