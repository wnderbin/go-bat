package highlight

import "os"

func ReadFile(file *os.File) string {
	content, _ := os.ReadFile(file.Name())
	return string(content)
}
