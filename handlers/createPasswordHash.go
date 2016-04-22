package hasherHandlers

import (
	"html/template"
	"net/http"

	"github.com/kkirsche/goPasswordHasher/models"
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
