package main

import (
	"fmt"
	"os"
)

func main() {
	fileUrl := os.Args[1]
	filePath := os.Args[2]
	fmt.Printf("will compare %s & %s ", fileUrl, filePath)
}
