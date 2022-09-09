package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/signin", signinHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/refresh", refreshHandler)

	http.ListenAndServe(":8080", nil)
}
