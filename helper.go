package avatar

// djb2
func StringHash(str string) int {
	hash := 5381
	for _, char := range str {
		hash = (hash << 5) + hash + int(char)
	}
	return hash
}
