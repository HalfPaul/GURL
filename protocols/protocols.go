package protocols

import (
	"GURL/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetData(url string, header string) (string, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the URL")
	}
	// TODO: Put this into function
	err = utils.HandleHeaders(header, request)
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

func PostData(url string, data string, header string) (string, error) {
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
	err = utils.HandleHeaders(header, request)
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

func PutData(url string, data string, header string) (string, error) {
	postBody, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Can't convert data to json")
	}
	fmt.Println(data)
	responseBody := bytes.NewBuffer(postBody)

	request, err := http.NewRequest(
		http.MethodPut,
		url,
		responseBody,
	)
	if err != nil {
		return "", fmt.Errorf("Could not resolve the URL")
	}

	err = utils.HandleHeaders(header, request)
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
