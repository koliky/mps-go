package app

import (
	"mps-go/pkg/view"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	if r.RequestURI != "/" {
		http.NotFound(w, r)
		return
	}
	view.Index(w, data)
}
