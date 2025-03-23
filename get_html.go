package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("request error status: %v", resp.StatusCode)
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		fmt.Println(resp.Header.Get("Content-Type"))
		return "", fmt.Errorf("invalid content type")
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(data), nil

}
