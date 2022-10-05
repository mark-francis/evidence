package app

import "fmt"

func Download(filename string) (string, error) {
	return fmt.Sprintf("Downloading %s...", filename), nil
}
