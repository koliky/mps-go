package app

import (
	"mps-go/pkg/view"
	"net/http"
)

func adminCreateUser(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	view.AdminCreateUser(w, data)
}

func adminListUser(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	search := r.URL.Query().Get("search")
	data["search"] = search
	view.AdminListUser(w, data)
}

func adminShowUser(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	userID := r.URL.Query().Get("id")
	if len(userID) != 0 {
		data["id"] = userID
		view.AdminShowUser(w, data)
	}
}

func adminUpdateUser(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	userID := r.URL.Query().Get("id")
	if len(userID) != 0 {
		data["id"] = userID
		view.AdminUpdateUser(w, data)
	}
}
