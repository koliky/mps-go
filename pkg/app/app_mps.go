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
	search := r.URL.Query().Get("search")
	data["search"] = search
	view.MpsCreateGroup(w, data)
}

func mpsCreatePart(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	search := r.URL.Query().Get("search")
	data["search"] = search
	groupID := r.URL.Query().Get("id")
	if len(groupID) != 0 {
		data["id"] = groupID
		view.MpsCreatePart(w, data)
	}
}
