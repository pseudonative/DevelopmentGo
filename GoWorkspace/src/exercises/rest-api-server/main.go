package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/items", handleItems)
	http.HandleFunc("/items", handleItem)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleItems(w http.ResponseWriter, r *http.Request) {

}

func handleItem(w http.ResponseWriter, r *http.Request) {

}
