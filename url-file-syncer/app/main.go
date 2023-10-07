package main

import (
	"os"
)

func main() {
	remoteFileUrl := os.Args[1]
	localFilePath := os.Args[2]

	message := "will compare " + remoteFileUrl + localFilePath
	setOutput("message", message)
}
