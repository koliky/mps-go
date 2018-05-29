package app

import (
	"mps-go/pkg/view"
	"net/http"
)

func mpsHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	view.MpsHome(w, data)
}

func mpsCreateForcast(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	view.MpsCreateForcast(w, data)
}

func mpsCreateGroup(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	view.MpsCreateGroup(w, data)
}
