package passwordhasher

import "net/http"

func init() {
	http.HandleFunc("/public/", PublicAssetsHandler)
	http.HandleFunc("/hash", HashPasswordHandler)
	http.HandleFunc("/", CreateHashHandler)
}
