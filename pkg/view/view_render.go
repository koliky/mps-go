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

// MpsCreatePart render view
func MpsCreatePart(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("mps/create-part.html"), w, data)
}

// AdminCreateUser render view
func AdminCreateUser(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("admin/create-user.html"), w, data)
}

// AdminListUser render view
func AdminListUser(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("admin/list-user.html"), w, data)
}

// AdminShowUser render view
func AdminShowUser(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("admin/show-user.html"), w, data)
}

// AdminUpdateUser render view
func AdminUpdateUser(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("admin/update-user.html"), w, data)
}
