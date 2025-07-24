package compressor

import "fmt"

func CmpressFolder(path string) (string, error) {
	return fmt.Sprintf("%s: Folder Zipped", path), nil
}
