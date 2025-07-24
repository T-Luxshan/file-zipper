package compressor

import "fmt"

func CompressFile(path string) (string, error) {
	return fmt.Sprintf("%s Zipped", path), nil
}
