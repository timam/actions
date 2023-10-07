package main

import (
	"io/ioutil"
)

func compareFiles(downloadedFile, localFile string) (bool, error) {
	downloadedContent, err := ioutil.ReadFile(downloadedFile)
	if err != nil {
		return false, err
	}
	localContent, err := ioutil.ReadFile(localFile)
	if err != nil {
		return false, err
	}
	return string(downloadedContent) == string(localContent), nil
}
