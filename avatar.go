package avatar

import (
	"fmt"
	"log"

	"github.com/swiftcarrot/color"
)

func GenerateGradient(str, text string, width, height int) (string, error) {
	hash := StringHash(str)

	log.Println(hash)

	r1, g1, b1 := color.HSL2RGB(float64(hash%360), 100, 50)
	r2, g2, b2 := color.HSL2RGB(float64((hash+120)%360), 100, 50)

	log.Println(r1, g1, b1)
	log.Println(r2, g2, b2)

	color1 := fmt.Sprintf("rgb(%.1f,%.1f,%.1f)", r1, g1, b1)
	color2 := fmt.Sprintf("rgb(%.1f,%.1f,%.1f)", r2, g2, b2)

	log.Println(color1, color2)

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

func GenerateSVG(username, text string, width, height int) (string, error) {
	avatar, err := GenerateGradient(username, text, width, height)
	if err != nil {
		return "", err
	}
	return avatar, nil
}
