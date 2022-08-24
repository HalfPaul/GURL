package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Result struct{}

var headerVar string
var protocolVar string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GURL",
	Short: "CURL alternative",
	Long:  `Command-line tool for transferring data using various network protocols.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		switch protocolVar {
		case "":
			fmt.Println(getData(args[0]))
		case "POST":
			fmt.Println("POST")
		case "PUT":
			fmt.Println("PUT")
		case "GET":
			fmt.Println(getData(args[0]))
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
}

func getData(url string) (string, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the URL")
	}

	if headerVar != "" {
		headerParts := strings.Split(headerVar, ": ")
		if len(headerParts) != 2 {
			return "", fmt.Errorf("Header is not set right.")
		}
		headerHead := headerParts[0]
		headerTail := headerParts[1]
		request.Header.Add(headerHead, headerTail)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the URL")
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the Response")
	}
	return string(responseBytes), nil
}
