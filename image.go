package avatar

import (
	"crypto/md5"
	"fmt"
	"io"
)

func generateGradient(username, text string, width, height int) (string, error) {
	hasher := md5.New()
	io.WriteString(hasher, username)
	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	color1 := hashStringToColor(hash)
	h, s, l := Hex2HSL(color1)

	s = s + s*0.5

	if l < 25 {
		l = l + l*3
	} else if l > 25 && l < 40 {
		l = l + l*0.8
	} else if l > 75 {
		l = l - l*0.4
	}

	color1 = HSL2Hex(h, s, l)
	color2 := getMatchingColor(h, s, l)

	avatar, err := CreateSVG(SVGData{
		Color1:   color1,
		Color2:   color2,
		Width:    width,
		Height:   height,
		Text:     text,
		FontSize: (float64(height) * 0.9) / float64(len(text)),
	})
	if err != nil {
		return "", err
	}
	return avatar, nil
}

func GenerateSVG(username, text string, width, height int) (string, error) {
	avatar, err := generateGradient(username, text, width, height)
	if err != nil {
		return "", err
	}
	return avatar, nil
}
