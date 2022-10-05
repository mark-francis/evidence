package app

import "fmt"

func Upload(filename string) (string, error) {
	return fmt.Sprintf("Uploading %s...", filename), nil
}
