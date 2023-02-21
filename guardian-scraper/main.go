package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var userAgents []string = []string{
	"Mozilla/5.0 (Windows; U; Windows NT 10.0;) AppleWebKit/535.9 (KHTML, like Gecko) Chrome/53.0.3673.116 Safari/537",
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_9_6; en-US) AppleWebKit/535.39 (KHTML, like Gecko) Chrome/53.0.1427.317 Safari/601",
	"Mozilla/5.0 (Android; Android 4.4.4; IQ4500 Quad Build/KOT49H) AppleWebKit/603.32 (KHTML, like Gecko)  Chrome/52.0.3108.336 Mobile Safari/603.2",
	"Mozilla/5.0 (Linux; U; Android 4.4.1; SM-G900V Build/KOT49H) AppleWebKit/600.42 (KHTML, like Gecko)  Chrome/50.0.1380.359 Mobile Safari/534.9",
	"Mozilla/5.0 (U; Linux x86_64) AppleWebKit/534.49 (KHTML, like Gecko) Chrome/51.0.3214.321 Safari/537",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 7_7_4; like Mac OS X) AppleWebKit/537.22 (KHTML, like Gecko)  Chrome/49.0.2963.192 Mobile Safari/603.6",
	"Mozilla/5.0 (U; Linux x86_64) Gecko/20100101 Firefox/45.1",
	"Mozilla/5.0 (Android; Android 4.4.4; LG-V500 Build/KOT49I) AppleWebKit/603.8 (KHTML, like Gecko)  Chrome/51.0.3635.362 Mobile Safari/600.3",
	"Mozilla/5.0 (Linux i684 x86_64; en-US) Gecko/20100101 Firefox/53.8",
	"Mozilla/5.0 (Android; Android 4.4.4; [HM NOTE|NOTE-III|NOTE2 1LTET) AppleWebKit/537.22 (KHTML, like Gecko)  Chrome/53.0.3786.254 Mobile Safari/601.2",
}
var tokens = make(chan struct{}, 5)

func randomUserAgent() string {
	return userAgents[rand.Intn(len(userAgents))]
}

func discoverLinks(response *http.Response, baseURL string) []string {
	if response != nil {
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			return nil
		}
		foundUrls := []string{}
		if doc != nil {
			doc.Find("a").Each(func(i int, s *goquery.Selection) {
				res, _ := s.Attr("href")
				foundUrls = append(foundUrls, res)
			})
		}
		return foundUrls
	}
	return nil
}

func getRequest(targetURL string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", targetURL, nil)
	req.Header.Set("User-Agent", randomUserAgent())
	return client.Do(req)
}

func checkRelative(href string, baseURL string) string {
	if strings.HasPrefix(href, "/") {
		return fmt.Sprintf("%s%s", baseURL, href)
	}
	return href
}

func resolveRelativeLinks(href string, baseURL string) (string, bool) {
	resultHref := checkRelative(href, baseURL)
	baseParse, _ := url.Parse(baseURL)
	resultParse, _ := url.Parse(resultHref)

	if baseParse != nil && resultParse != nil {
		if baseParse.Host == resultParse.Host {
			return resultHref, true
		} else {
			return "", false
		}
	}
	return "", false
}

// func parseHTML(response *http.Response) {
// 	//  Extract HTML to get data
// }

func crawl(targetURL string, baseURL string) []string {
	fmt.Println(targetURL)
	tokens <- struct{}{}
	response, err := getRequest(targetURL)
	<-tokens
	if err != nil {
		return nil
	}
	links := discoverLinks(response, baseURL)
	foundUrls := []string{}

	for _, link := range links {
		absoluteLink, ok := resolveRelativeLinks(link, baseURL)
		if ok {
			if absoluteLink != "" {
				foundUrls = append(foundUrls, absoluteLink)
			}
		}
	}
	return foundUrls
}

func main() {
	worklist := make(chan []string)
	var n int = 1
	baseDomain := "https://www.theguardian.com"

	go func() {
		worklist <- []string{baseDomain}
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist

		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, baseURL string) {
					foundLinks := crawl(link, baseURL)
					if foundLinks != nil {
						worklist <- foundLinks
					}
				}(link, baseDomain)
			}
		}
	}
}
