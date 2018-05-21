package app

import (
	"mps-go/pkg/view"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	view.Login(w, nil)
}
