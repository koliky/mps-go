package app

import (
	"mps-go/pkg/view"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	view.Login(w, data)
}
