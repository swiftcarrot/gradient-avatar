package main

import (
	// "crypto/md5"
	// "encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// func generateGradient(username, text string, width, height int) {
// 	hasher := md5.New()
// 	hasher.Write([]byte(text))
// 	hash := hex.EncodeToString(hasher.Sum(nil))
// }

func generateAvatar(name string) string {
	avatar := `<?xml version="1.0" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg width="$WIDTH" height="$HEIGHT" viewBox="0 0 $WIDTH $HEIGHT" version="1.1" xmlns="http://www.w3.org/2000/svg">
  <g>
    <defs>
      <linearGradient id="avatar" x1="0" y1="0" x2="1" y2="1">
        <stop offset="0%" stop-color="$FIRST"/>
        <stop offset="100%" stop-color="$SECOND"/>
      </linearGradient>
    </defs>
    <rect fill="url(#avatar)" x="0" y="0" width="$WIDTH" height="$HEIGHT"/>
    <text x="50%" y="50%" alignment-baseline="central" dominant-baseline="middle" text-anchor="middle" fill="#fff" font-family="sans-serif" font-size="$FONTSIZE">$TEXT</text>
  </g>
</svg>
`

	avatar = strings.Replace(avatar, "$WIDTH", "100", -1)
	avatar = strings.Replace(avatar, "$HEIGHT", "100", -1)
	// todo: color for username
	avatar = strings.Replace(avatar, "$FIRST", "#5DC3FF", -1)
	avatar = strings.Replace(avatar, "$SECOND", "#A8FF63", -1)
	avatar = strings.Replace(avatar, "$TEXT", string([]rune(name)[0]), -1)
	avatar = strings.Replace(avatar, "$FONTSIZE", "48", -1)

	return avatar
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "max-age=2592000, public")

	fmt.Fprint(w, generateAvatar("王"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
