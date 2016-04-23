package passwordhasher

import "net/http"

func init() {
	http.HandleFunc("/hash", HashPasswordHandler)
	http.HandleFunc("/", CreateHashHandler)
}
