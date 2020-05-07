package handlers

import (
	"fmt"
	"net/http"
)

//DefaultHandler ...
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello world!")

}
