package main

import (
	"fmt"
	"net/http"
)

func main() {

	returnFailure := false

	// Request fails every other time to rest fallback cache
	http.HandleFunc("/failodd", func(w http.ResponseWriter, r *http.Request) {
		if returnFailure {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Uh oh something bad happened")
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Everything is fine")
		}

		returnFailure = !returnFailure
	})

	//TODO: javacript with fingerprint example

	http.ListenAndServe(":8000", nil)
}
