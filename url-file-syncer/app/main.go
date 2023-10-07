package main

import (
	"os"
)

func main() {
	remoteFileUrl := os.Args[1]
	//localFilePath := os.Args[2]

	file, err := downloadFile(remoteFileUrl)
	if err != nil {
		message := "error when downloading file " + file
		setOutput("message", message)
		return
	} else {
		message := "successfully downloaded file " + file
		setOutput("message", message)
		return
	}

}
