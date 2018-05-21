package app

import (
	"mps-go/pkg/view"
	"net/http"
)

func mpsHome(w http.ResponseWriter, r *http.Request) {
	view.MpsHome(w, nil)
}

func mpsCreateForcast(w http.ResponseWriter, r *http.Request) {
	view.MpsCreateForcast(w, nil)
}

func mpsCreateGroup(w http.ResponseWriter, r *http.Request) {
	view.MpsCreateGroup(w, nil)
}
