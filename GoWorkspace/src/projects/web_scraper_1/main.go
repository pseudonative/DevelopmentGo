package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.warmvalleyhc.com/"
	content, err := fetchURL(url)
	if err != nil {
		fmt.Println("Error fetching URL: ", err)
		return
	}
	// fmt.Println(content)
	parseHTML(content)
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

func parseHTML(htmlContent string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println("Error creating document: ", err)
		return
	}
	doc.Find("title").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		fmt.Printf("Title %d: %s\n", index+1, title)
	})
}
