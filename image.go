package avatar

import (
	"crypto/md5"
	"fmt"
	"io"
)

func generateGradient(username, text string, width, height int) (string, error) {
	h := md5.New()
	io.WriteString(h, username)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	firstColor := hashStringToColor(hash)
	// log.Println(firstColor)
	rgb, err := HTMLToRGB(firstColor)
	if err != nil {
		return "", err
	}

	// log.Println(rgb)
	hsl := rgb.ToHSL()
	// log.Println(hsl)

	hsl = HSL{
		H: hsl.H,
		S: hsl.S * 0.5,
		L: hsl.L,
	}

	if hsl.L < 25 {
	} else if hsl.L > 25 && hsl.L < 40 {
	} else if hsl.L > 75 {
	}

	firstColor = hsl.ToHTML()
	secondColor := firstColor

	avatar, err := CreateSVG(SVGData{
		First:    firstColor,
		Second:   secondColor,
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
