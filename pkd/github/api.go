package github

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gritzkoo/golang-health-checker-lw/pkg/healthchecker"
)

// Status is a function to be used to healthchecker package test integration!
func Status() healthchecker.CheckResponse {
	result := healthchecker.CheckResponse{
		URL: "https://github.com/status",
	}
	client := http.Client{
		Timeout: time.Duration(3 * time.Second),
	}
	request, err := http.NewRequest("GET", result.URL, nil)
	if err != nil {
		result.Error = err
		return result
	}

	response, err := client.Do(request)

	if err != nil {
		result.Error = err
		return result
	}
	if response.StatusCode != http.StatusOK {
		result.Error = fmt.Errorf("Error call github.com/status get code: %d", response.StatusCode)
	}
	return result

}
