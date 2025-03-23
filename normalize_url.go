package main

import (
	"net/url"
)

func normalizeURL(input string) (string, error) {
	url, err := url.Parse(input)

	if err != nil {
		return "", err
	}

	result := url.Host + url.Path

	if len(url.RawQuery) > 0 {
		result = result + "?" + url.RawQuery
	}

	if len(url.RawFragment) > 0 {
		result = result + "#" + url.RawFragment
	}

	return result, nil
}
