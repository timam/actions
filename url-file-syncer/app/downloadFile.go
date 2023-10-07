package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadFile(url string) error {
	fileName := filepath.Base(url)

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
