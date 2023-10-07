package main

import (
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	remoteFileUrl := os.Args[1]
	localFilePath := filepath.Join(os.Getenv("GITHUB_WORKSPACE"), os.Args[2])

	remoteFilePath, err := downloadFile(remoteFileUrl)

	if err != nil {
		message := "error when downloading file " + remoteFilePath
		setOutput("message", message)
		return
	} else {
		result, err := compareFiles(remoteFilePath, localFilePath)
		if err != nil {
			message := "error when comparing files. result :" + strconv.FormatBool(result) + " error: " + err.Error()
			setOutput("message", message)
			return
		} else {
			message := "file compare success. result : " + strconv.FormatBool(result)
			setOutput("message", message)
			return
		}
	}

}
