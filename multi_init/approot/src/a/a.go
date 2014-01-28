package a

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
}

func UsedExternally() {
}
