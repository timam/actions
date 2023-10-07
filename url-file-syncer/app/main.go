package main

import (
	"fmt"
	"os"
)

func main() {
	remoteFileUrl := os.Args[1]
	localFilePath := os.Args[2]

	err := downloadFile(remoteFileUrl)
	if err != nil {
		fmt.Printf("Error while doenloading file %s", err)
	}

	message := "will compare " + remoteFileUrl + " " + localFilePath
	setOutput("message", message)
}
