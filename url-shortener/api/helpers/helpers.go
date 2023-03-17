package helpers

import (
	"net/url"
	"os"
	"strings"
)

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	// http(s)://(www.)localhost
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	return newURL != os.Getenv("DOMAIN")
}

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func IsValidURL(urlString string) bool {
	parseUrl, err := url.Parse(urlString)
	if err != nil {
		return false
	}
	return parseUrl.Scheme == "" || parseUrl.Scheme == "http" || parseUrl.Scheme == "https"
}
