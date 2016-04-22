package hasherHandlers

import (
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile(`^/(css|js|img)/([a-zA-Z0-9\.\_]+)$`)

// PublicAssets serves publicly accessible web content
func PublicAssets(w http.ResponseWriter, r *http.Request) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
