package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleHeaders(header string, request *http.Request) error {
	if header != "" {
		headerParts := strings.Split(header, ": ")
		if len(headerParts) != 2 {
			return fmt.Errorf("Header is not set right.")
		}
		headerHead := headerParts[0]
		headerTail := headerParts[1]
		request.Header.Add(headerHead, headerTail)
	}

}
