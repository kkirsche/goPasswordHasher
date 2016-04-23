package passwordhasher

import (
	"html/template"
	"net/http"
)

const (
	noPassword = "No password was provided. Please enter a password and try again."
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
	switch password {
	case "":
		p.SHA512 = noPassword
		p.SHA256 = noPassword
		p.Bcrypt = noPassword
		p.APR1 = noPassword
		p.MD5 = noPassword
	default:
		p.SHA512, err = GenerateSHA512FromString(password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		p.SHA256, err = GenerateSHA256FromString(password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		p.Bcrypt, err = GenerateBcryptFromString(password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		p.APR1, err = GenerateAPR1FromString(password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		p.MD5, err = GenerateMD5FromString(password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	err = templates.ExecuteTemplate(w, "hashedPassword.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
