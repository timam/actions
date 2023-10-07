package main

import (
	"fmt"
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

func syncLocalFileWithRemote(downloadedFile, localFile string) error {
	downloaded, err := os.Open(downloadedFile)
	if err != nil {
		return fmt.Errorf("failed to open downloaded file: %w", err)
	}
	defer downloaded.Close()

	local, err := os.OpenFile(localFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to create or open local file: %w", err)
	}
	defer local.Close()

	_, err = io.Copy(local, downloaded)
	if err != nil {
		return fmt.Errorf("failed to sync local file with remote: %w", err)
	}

	return nil
}

func deleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
