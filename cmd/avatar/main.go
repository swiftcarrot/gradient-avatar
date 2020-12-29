package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swiftcarrot/avatar"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "max-age=2592000, public")

	text := r.URL.Query().Get("text")
	username := r.URL.Query().Get("username")
	svg, err := avatar.GenerateSVG(username, text, 100, 100)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, svg)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
