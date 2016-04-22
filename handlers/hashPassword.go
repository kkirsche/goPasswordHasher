package hasherHandlers

import (
	"net/http"

	"github.com/kkirsche/ansiblePasswordGenerator/hasher"
	"github.com/kkirsche/ansiblePasswordGenerator/models"
)

// HashPasswordHandler is responsible for handling requests to hash a password
// from the Create Password Hash website.
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
