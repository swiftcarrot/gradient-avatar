package avatar

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/swiftcarrot/color"
)

func GenerateGradient(str, text string, width, height int) (string, error) {
	hasher := md5.New()
	io.WriteString(hasher, str)
	hash := fmt.Sprintf("%x", hasher.Sum(nil))
	color1 := hashStringToColor(hash)
	h, s, l := color.Hex2HSL(color1)

	s = s + s*0.5
	if l < 25 {
		l = l + l*2.5
	} else if l > 25 && l < 40 {
		l = l + l*0.8
	} else if l > 75 {
		l = l - l*0.4
	}

	color1 = color.HSL2Hex(h, s, l)
	color2 := getMatchingColor(h, s, l)

	avatar, err := CreateSVG(SVGData{
		Color1:   color1,
		Color2:   color2,
		Width:    width,
		Height:   height,
		Text:     text,
		FontSize: (float64(height) * 0.5) / float64(len([]rune(text))),
	})
	if err != nil {
		return "", err
	}
	return avatar, nil
}

func GenerateSVG(str, text string, width, height int) (string, error) {
	avatar, err := GenerateGradient(str, text, width, height)
	if err != nil {
		return "", err
	}
	return avatar, nil
}
