package handlers

import (
	"fmt"
	"net/http"
)

// Home is the home page handler.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}
