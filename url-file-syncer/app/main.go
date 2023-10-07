package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	remoteFileUrl := os.Args[1]
	localFilePath := os.Args[2]

	downloadedFileName, err := downloadFile(remoteFileUrl)
	if err != nil {
		fmt.Printf("Error while doenloading file %s", err)
	}

	files, err := compareFiles(downloadedFileName, localFilePath)
	if err != nil {
		return
	}

	//message := "will compare " + remoteFileUrl + " " + localFilePath
	message := "result of compare is " + strconv.FormatBool(files)
	setOutput("message", message)
}
