package handlers

import (
	"fmt"
	"net/http"
)

func handleItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "This will show a list of items.")
	} else {
		http.Error(w, "Method is onotsupported.", http.StatusMethodNotAllowed)
	}
}

func handleItem(w http.ResponseWriter, r *http.Request) {

}
