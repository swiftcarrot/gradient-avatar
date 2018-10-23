package avatar

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

func GenerateGradient(username, text string, width, height int) (string, error) {
	hasher := md5.New()
	io.WriteString(hasher, username)
	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	color1 := hashStringToColor(hash)

	log.Println("color1", color1)

	h, s, l := Hex2HSL(color1)

	log.Println("hsl", h, s, l)

	s = s + s*0.5

	if l < 25 {
		l = l + l*3
	} else if l > 25 && l < 40 {
		l = l + l*0.8
	} else if l > 75 {
		l = l - l*0.4
	}

	log.Println("hsl", h, s, l)

	color1 = HSL2Hex(h, s, l)
	color2 := getMatchingColor(h, s, l)

	log.Println(color1, color2)

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
	avatar, err := GenerateGradient(username, text, width, height)
	if err != nil {
		return "", err
	}
	return avatar, nil
}
