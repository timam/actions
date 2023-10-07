package main

import (
	"os"
)

func main() {
	fileUrl := os.Args[1]
	filePath := os.Args[2]
	message := "will compare " + fileUrl + " " + filePath
	setOutput(message)
}
