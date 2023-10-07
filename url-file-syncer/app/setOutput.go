package main

import (
	"fmt"
	"os"
)

func setOutput(id string, output string) {
	outputFile := os.Getenv("GITHUB_OUTPUT")
	if outputFile == "" {
		fmt.Println("GITHUB_OUTPUT environment variable not set")
		os.Exit(1)
	}

	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s=%s\n", id, output)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}
