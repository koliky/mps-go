package app

import (
	"mps-go/pkg/view"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		http.NotFound(w, r)
		return
	}
	view.Index(w, nil)
}
