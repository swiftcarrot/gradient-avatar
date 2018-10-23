package avatar

import (
	"fmt"
	"testing"
)

func TestCreateSVG(t *testing.T) {
	svg, err := CreateSVG(SVGData{
		First:    "#3498db",
		Second:   "#889912",
		Width:    30,
		Height:   30,
		Text:     "WZ",
		FontSize: 20,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(svg)
}
