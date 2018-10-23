package avatar

import (
	"fmt"
	"testing"
)

func TestHex2RGB(t *testing.T) {}

func TestHSL2Hex(t *testing.T) {
	hex := HSL2Hex(238, 116, 45)
	fmt.Println(hex)
}

func TestHSV2RGB(t *testing.T) {
	r, g, b := HSV2RGB(4, 59, 93)
	fmt.Println("HSV2RGB", r, g, b)
}
