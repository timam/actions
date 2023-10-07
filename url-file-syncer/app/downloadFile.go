package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadFile(url string) (string, error) {
	fileName := filepath.Base(url)
	filePath := filepath.Join(os.Getenv("GITHUB_WORKSPACE"), fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
