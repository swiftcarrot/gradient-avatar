package avatar

import (
	"fmt"
	"math"
)

func Hex2RGB(hex string) (float64, float64, float64) {
	var r, g, b byte
	fmt.Sscanf(hex, "#%2x%2x%2x", &r, &g, &b)
	return float64(r), float64(g), float64(b)
}

func Hex2HSL(hex string) (float64, float64, float64) {
	r, g, b := Hex2RGB(hex)
	h, s, v := RGB2HSV(r, g, b)
	h, s, l := HSV2HSL(h, s, v)
	return h, s, l
}

func RGB2HSV(r, g, b float64) (float64, float64, float64) {
	var h, s, v float64
	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	delta := max - min

	if delta == 0 {
		h = 0
	} else if r == max {
		h = math.Mod((g-b)/delta, 6.0)
	} else if g == max {
		h = (b-r)/delta + 2
	} else {
		h = (r-g)/delta + 4
	}

	h = math.Round(h * 60)
	if h < 0 {
		h += 360
	}

	if max == 0 {
		s = 0
	} else {
		s = math.Round((delta / max) * 100)
	}

	v = math.Round((max / 255) * 100)

	return h, s, v
}

func HSL2HSV(h, s, l float64) (float64, float64, float64) {
	if l < 50 {
		s = s * l / 100
	} else {
		s = s * (100 - l) / 100
	}
	return h, ((2 * s) / (l + s)) * 100, l + s
}

func HSV2HSL(h, s, v float64) (float64, float64, float64) {
	hh := ((200 - s) * v) / 100

	if hh < 100 {
		s = (s * v) / hh
	} else {
		s = (s * v) / (200 - hh)
	}
	l := hh / 2
	return h, s, l
}

func RGB2Hex(r, g, b float64) string {
	return fmt.Sprintf("#%02x%02x%02x", byte(r), byte(g), byte(b))
}

func HSV2RGB(h, s, v float64) (float64, float64, float64) {
	s = s / 100
	v = v / 100
	rgb := []float64{}

	c := v * s
	hh := h / 60
	x := c * (1 - math.Abs(math.Mod(hh, 2)-1))
	m := v - c

	switch int(hh) {
	case 0:
		rgb = []float64{c, x, 0}

	case 1:
		rgb = []float64{x, c, 0}
	case 2:
		rgb = []float64{0, c, x}

	case 3:
		rgb = []float64{0, x, c}

	case 4:
		rgb = []float64{x, 0, c}

	case 5:
		rgb = []float64{c, 0, x}
	}

	return math.Round(255 * (rgb[0] + m)), math.Round(255 * (rgb[0] + m)), math.Round(255 * (rgb[0] + m))
}

func HSV2Hex(h, s, v float64) string {
	r, g, b := HSV2RGB(h, s, v)
	return RGB2Hex(r, g, b)
}

func HSL2Hex(h, s, l float64) string {
	h, s, v := HSL2HSV(h, s, l)
	return HSV2Hex(h, s, v)
}

func HSL2RGB(h, s, l float64) (float64, float64, float64) {
	h, s, v := HSL2HSV(h, s, l)
	r, g, b := HSV2RGB(h, s, v)

	return r, g, b
}
