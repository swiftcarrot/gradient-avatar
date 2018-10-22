package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wangzuo/avatar"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "max-age=2592000, public")

	svg, err := avatar.GenerateSVG("wangzuo", "wz", 100, 100)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, svg)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}