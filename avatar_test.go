package avatar

import (
	"testing"
)

func TestGenerateGradient(t *testing.T) {
	_, err := GenerateGradient("swiftcarrot", "王", 100, 100)
	if err != nil {
		t.Fatal(err)
	}
}
