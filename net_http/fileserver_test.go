package net_http

import (
	"net/http"
	"testing"
)

func Test_Http_FileServer(t *testing.T) {
	t.Log(http.ListenAndServe(":8080",
		http.FileServer(
			http.Dir(`e:\\ProjectResources\\GoLangProjects\Go-Standard-Library`),
		)))
}
