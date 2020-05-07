package utils

import (
	"fmt"
	"net/http"
)

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "encountered error: %+v\n", err)
}

// RespondWithSuccess ...
func RespondWithSuccess(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%+v\n", message)
}
