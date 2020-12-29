package avatar

import (
	"testing"
)

func TestCreateSVG(t *testing.T) {
	_, err := CreateSVG(SVGData{
		Color1:   "#3498db",
		Color2:   "#889912",
		Width:    30,
		Height:   30,
		Text:     "WZ",
		FontSize: 20,
	})
	if err != nil {
		t.Fatal(err)
	}
}
