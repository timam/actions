package main

import (
	"os"
)

func compareFiles(downloadedFile, localFile string) (bool, error) {
	downloadedContent, err := os.ReadFile(downloadedFile)
	if err != nil {
		return false, err
	}
	localContent, err := os.ReadFile(localFile)
	if err != nil {
		return false, err
	}
	return string(downloadedContent) == string(localContent), nil
}
