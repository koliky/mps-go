package app

import (
	"net/http"
)

// CreateRouter is make all router in app
func CreateRouter(mux *http.ServeMux) {
	mux.Handle("/-/", http.StripPrefix("/-", http.FileServer(http.Dir("static"))))

	midelwareMux := http.NewServeMux()
	midelwareMux.HandleFunc("/", index)
	midelwareMux.HandleFunc("/mps", mpsHome)
	midelwareMux.HandleFunc("/mps/createforcast", mpsCreateForcast)
	midelwareMux.HandleFunc("/mps/creategroup", mpsCreateGroup)
	midelwareMux.HandleFunc("/login", login)

	mux.Handle("/", midelware(midelwareMux))
}

func midelware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}