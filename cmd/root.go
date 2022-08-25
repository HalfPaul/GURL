package cmd

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"

	"GURL/protocols"
)

type Result struct{}

var headerVar string
var protocolVar string
var dataVar string

// TODO : Put color into all print statements.
var rootCmd = &cobra.Command{
	Use:   "GURL",
	Short: "CURL alternative",
	Long:  `Command-line tool for transferring data using various network protocols.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		switch protocolVar {
		case "":
			fmt.Println(protocols.GetData(args[0], headerVar))
		case "POST":
			if dataVar == "" {
				fmt.Println("No data was given for POST request.")
			}
			fmt.Println(protocols.PostData(args[0], dataVar, headerVar))
		case "PUT":
			if dataVar == "" {
				fmt.Println("No data was given for PUT request.")
			}
			fmt.Println(protocols.PutData(args[0], dataVar, headerVar))
		case "GET":
			fmt.Println(protocols.GetData(args[0], headerVar))
		default:
			fmt.Println(fmt.Sprintf("Protocol by the name %s doesn't exist", protocolVar))
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
