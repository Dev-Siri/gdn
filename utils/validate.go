package utils

import "net/url"

func IsValidURL(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)

	return err == nil
}
