package view

import "net/http"

// Index render view
func Index(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("index.html"), w, data)
}

// Login render view
func Login(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("login.html"), w, data)
}

// MpsHome render view
func MpsHome(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("mps/home.html"), w, data)
}

// MpsCreateForcast render view
func MpsCreateForcast(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("mps/create-forcast.html"), w, data)
}

// MpsCreateGroup render view
func MpsCreateGroup(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("mps/create-group.html"), w, data)
}
