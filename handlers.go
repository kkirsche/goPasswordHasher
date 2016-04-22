package main

import (
	"html/template"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("tmpl/createHash.html", "tmpl/hashedPassword.html"))

// CreateHashHandler handles serving the create password hash portion of this
// website.
func CreateHashHandler(w http.ResponseWriter, r *http.Request) {
	p := &CreateHash{
		Title: "Password Hasher",
	}
	err := templates.ExecuteTemplate(w, "createHash.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HashPasswordHandler is responsible for handling requests to hash a password
// from the Create Password Hash website.
func HashPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	p := &HashedPassword{
		Title: "Password Hasher",
	}

	password := r.FormValue("password")
	hashType := r.FormValue("hashType")
	p.PasswordHash, err = HashPassword(password, hashType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = templates.ExecuteTemplate(w, "hashedPassword.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validAssetsPath = regexp.MustCompile(`^/(css|js|img)/([a-zA-Z0-9\.\_]+)$`)

// PublicAssetsHandler serves publicly accessible web content
func PublicAssetsHandler(w http.ResponseWriter, r *http.Request) {
	m := validAssetsPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
