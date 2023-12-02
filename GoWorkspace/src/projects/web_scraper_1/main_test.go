package main

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestParseHTML(t *testing.T) {
	htmlContent := `<html lang="en">
	<head>
	  <meta charset="UTF-8" />
	  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	  <title>🏘 Warm Valley HC, LLC</title>
	  <link rel="stylesheet" href="/css/styles.css">
	</head>`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		t.Errorf("Error creating document from reader: %v", err)
	}

	found := false
	doc.Find("title").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		if title == "🏘 Warm Valley HC, LLC" {
			found = true
		}
	})

	if !found {
		t.Errorf("Failed to find the title in the HTML content")
	}
}
