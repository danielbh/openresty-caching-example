package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var fakeVersion = "asdllkdzs"

func main() {
	returnFailure := false

	rand.Seed(time.Now().UnixNano())

	// Don't block main goroutine
	// change fake version every 2 seconds
	go repeat(2000*time.Millisecond, func() {
		newFakeVersion := ""
		for i := 1; i <= 10; i++ {
			newFakeVersion += string(rand.Intn(122-97) + 97)
		}
		fakeVersion = newFakeVersion
	})

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

	http.ListenAndServe(":8000", nil)
}

func repeat(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}
