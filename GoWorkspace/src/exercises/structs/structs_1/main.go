package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	myBook := Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan and Brian W. Kernington",
		Pages:  380,
	}
	PrintBookInfo(myBook)
}

func PrintBookInfo(b Book) {
	fmt.Printf("Book Title: %s\n Book Author: %s\n Pages: %d\n", b.Title, b.Author, b.Pages)
}

func UpdateBookPages(b *Book, newPages int)  {
	b.Pages = newPages
}