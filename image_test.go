package avatar

import (
	"log"
	"testing"
)

func TestGenerateGradient(t *testing.T) {
	avatar, err := GenerateGradient("wangzuo", "王", 100, 100)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(avatar)
}
