package cmd

import (
	"GURL/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
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
			fmt.Println(getData(args[0]))
		case "POST":
			if dataVar == "" {
				fmt.Println("No data was given for POST request.")
			}
			fmt.Println(postData(args[0], dataVar))
		case "PUT":
			if dataVar == "" {
				fmt.Println("No data was given for PUT request.")
			}
			fmt.Println(putData(args[0], dataVar))
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
	rootCmd.Flags().StringVarP(&dataVar, "data", "d", "", "Variable for data.")
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
	// TODO: Put this into function
	err = utils.HandleHeaders(headerVar, request)
	if err != nil {
		return "", err
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

func postData(url string, data string) (string, error) {
	postBody, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Can't convert data to json")
	}
	fmt.Println(data)
	responseBody := bytes.NewBuffer(postBody)

	request, err := http.NewRequest(
		http.MethodPost,
		url,
		responseBody,
	)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the URL")
	}
	err = utils.HandleHeaders(headerVar, request)
	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the Response")
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the Response")
	}
	return string(responseBytes), nil
}

func putData(url string, data string) (string, error) {
	postBody, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Can't convert data to json")
	}
	fmt.Println(data)
	responseBody := bytes.NewBuffer(postBody)

	request, err := http.NewRequest(
		http.MethodPost,
		url,
		responseBody,
	)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the URL")
	}

	err = utils.HandleHeaders(headerVar, request)
	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the Response")
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the Response")
	}
	return string(responseBytes), nil
}
