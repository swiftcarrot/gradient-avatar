package avatar

import "fmt"

func djb2(str string) int {
	hash := 5381
	for _, char := range str {
		hash = (hash << 5) + hash + int(char)
	}
	return hash
}

func hashStringToColor(str string) string {
	hash := djb2(str)

	ir := (hash & 0xff0000) >> 16
	ig := ((hash & 0x00ff00) >> 8)
	ib := hash & 0x0000ff

	r := fmt.Sprintf("0%x", ir)
	g := fmt.Sprintf("0%x", ig)
	b := fmt.Sprintf("0%x", ib)

	r = r[len(r)-2:]
	g = g[len(g)-2:]
	b = b[len(b)-2:]

	return fmt.Sprintf("#%s%s%s", r, g, b)
}
