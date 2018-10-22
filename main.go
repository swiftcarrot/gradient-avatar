package main // import github.com/wangzuo/avatar

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strings"
)

type RGB struct {
	R, G, B float64
}

type HSL struct {
	H, S, L float64
}

func HTMLToRGB(in string) (RGB, error) {
	if in[0] == '#' {
		in = in[1:]
	}

	if len(in) != 6 {
		return RGB{}, errors.New("Invalid string length")
	}

	var r, g, b byte
	if n, err := fmt.Sscanf(in, "%2x%2x%2x", &r, &g, &b); err != nil || n != 3 {
		return RGB{}, err
	}

	return RGB{float64(r) / 255, float64(g) / 255, float64(b) / 255}, nil
}

func (c RGB) ToHSL() HSL {
	var h, s, l float64

	r := c.R
	g := c.G
	b := c.B

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	l = (max + min) / 2

	delta := max - min
	if delta == 0 {

		return HSL{0, 0, l}
	}

	if l < 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
	}

	r2 := (((max - r) / 6) + (delta / 2)) / delta
	g2 := (((max - g) / 6) + (delta / 2)) / delta
	b2 := (((max - b) / 6) + (delta / 2)) / delta
	switch {
	case r == max:
		h = b2 - g2
	case g == max:
		h = (1.0 / 3.0) + r2 - b2
	case b == max:
		h = (2.0 / 3.0) + g2 - r2
	}

	switch {
	case h < 0:
		h += 1
	case h > 1:
		h -= 1
	}

	return HSL{h, s, l}
}

func generateGradient(username, text string, width, height int) {
	h := md5.New()
	io.WriteString(h, username)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	firstColor := hashStringToColor(hash)
	log.Println(firstColor)
	rgb, err := HTMLToRGB(firstColor)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rgb)

	log.Println(firstColor)
}

func hashStringToColor(str string) string {
	hash := djb2(str)

	ir := (hash & 0xff0000) >> 16
	ig := ((hash & 0x00ff00) >> 8)
	ib := hash & 0x0000ff

	r := fmt.Sprintf("0%x", ir)
	g := fmt.Sprintf("0%x", ig)
	b := fmt.Sprintf("0%x", ib)

	r = r[len(r)-2:]
	g = g[len(g)-2:]
	b = b[len(b)-2:]

	return fmt.Sprintf("#%s%s%s", r, g, b)
}

func djb2(str string) int {
	hash := 5381
	for _, char := range str {
		hash = (hash << 5) + hash + int(char)
	}
	return hash
}

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
	// log.Println(djb2("你好"))
	// log.Println(hashStringToColor("你好"))

	generateGradient("你好", "test", 100, 100)

	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
