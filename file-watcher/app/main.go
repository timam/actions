package main

import (
	"fmt"
	"os"
)

func main() {
	fileUrl := os.Args[1]
	filePath := os.Args[2]
	message := "will compare" + fileUrl + " " + filePath
	setOutput(message)
}

func setOutput(message string) {
	outputFile := os.Getenv("GITHUB_OUTPUT")
	if outputFile == "" {
		fmt.Println("GITHUB_OUTPUT environment variable not set")
		os.Exit(1)
	}

	// Open the file for appending
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Write the message to the file
	_, err = fmt.Fprintf(file, "message=%s\n", message)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}
