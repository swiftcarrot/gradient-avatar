package avatar

import (
	"fmt"
	"math"

	"github.com/swiftcarrot/color"
)

func djb2(str string) int {
	hash := 5381
	for _, char := range str {
		hash = (hash << 5) + hash + int(char)
	}
	return hash
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

func shouldChangeColor(r, g, b float64) bool {
	val := 765 - (r + g + b)
	return val < 250 || val > 700
}

func getMatchingColor(h, s, l float64) string {
	r, g, b := color.HSL2RGB(h, s, l)
	yiq := (r*299 + g*587 + b*114) / 1000

	if yiq < 128 {
		s = s + s*0.3
	} else {
		s = s - s*0.3
	}

	h = math.Mod((h + 90), 360)
	if h < 0 {
		h = 360 + h
	}

	r, g, b = color.HSL2RGB(h, s, l)
	if shouldChangeColor(r, g, b) {
		h = math.Mod((h - 200), 360)
		if h < 0 {
			h = 360 + h
		}

		s = s + s*0.5
	}

	return color.HSL2Hex(h, s, l)
}
