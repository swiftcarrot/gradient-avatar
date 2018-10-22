package avatar // import github.com/wangzuo/avatar

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "max-age=2592000, public")

	avatar, err := GenerateSVG("wangzuo", "wz", 100, 100)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, avatar)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
