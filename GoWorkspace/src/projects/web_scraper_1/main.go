package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PageInfo struct {
	Title string
}

func main() {
	var url string
	flag.StringVar(&url, "url", "https://www.warmvalleyhc.com", "URL to scrape")
	flag.Parse()

	content, err := fetchURL(url)
	if err != nil {
		fmt.Println("Error fetching URL: ", err)
		return
	}

	pageInfo := parseHTML(content)
	jsonData, err := json.MarshalIndent(pageInfo, "", " ")
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}
	fmt.Println(string(jsonData))

}

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func parseHTML(htmlContent string) PageInfo {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println("Error creating document: ", err)
		return PageInfo{}
	}

	pageInfo := PageInfo{}

	title := doc.Find("title").First().Text()
	pageInfo.Title = title

	doc.Find("title").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		fmt.Printf("Title %d: %s\n", index+1, title)
	})
	return pageInfo
}
