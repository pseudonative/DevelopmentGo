package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/items", handleItems)
	http.HandleFunc("/item", handleItem)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
