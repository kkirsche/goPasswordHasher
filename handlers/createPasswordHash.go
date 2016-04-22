package hasherHandlers

import (
	"html/template"
	"net/http"

	"github.com/kkirsche/ansiblePasswordGenerator/hasher"
	"github.com/kkirsche/ansiblePasswordGenerator/models"
)

var templates = template.Must(template.ParseFiles("tmpl/createHash.html", "tmpl/hashedPassword.html"))

// CreateHashHandler handles serving the create password hash portion of this
// website.
func CreateHashHandler(w http.ResponseWriter, r *http.Request) {
	p := &hasherModels.CreateHash{
		Title: "Create SHA-512 Password Hash",
	}
	err := templates.ExecuteTemplate(w, "createHash.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HashPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	p := &hasherModels.HashedPassword{
		Title: "SHA-512 Password Hash",
	}
	password := r.FormValue("password")
	if password != "" {
		p.PasswordHash, err = passwordHasher.GenerateSHA512FromString(password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		p.PasswordHash = "Password cannot be a blank value. Please try again."
	}
	err = templates.ExecuteTemplate(w, "hashedPassword.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
