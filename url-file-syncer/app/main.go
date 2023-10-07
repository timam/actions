package main

import (
	"os"
	"path/filepath"
)

func main() {
	remoteFileUrl := os.Args[1]
	localFilePath := filepath.Join(os.Getenv("GITHUB_WORKSPACE"), os.Args[2])

	remoteFilePath, err := downloadFile(remoteFileUrl)

	if err != nil {
		message := "error when downloading file " + remoteFilePath + err.Error()
		setOutput("message", message)
	} else {
		alreadyUpDated, err := compareFiles(remoteFilePath, localFilePath)
		if err != nil {
			message := "error while comparing files - " + err.Error()
			setOutput("message", message)
		}

		if alreadyUpDated == true {
			message := "files are already in sync, nothing to do"
			setOutput("message", message)
		} else {
			err := syncLocalFileWithRemote(remoteFilePath, localFilePath)
			if err != nil {
				message := "error while syncing local file compared to remote " + err.Error()
				setOutput("message", message)
			} else {
				message := "successfully synced local file with remote"
				setOutput("message", message)
			}

		}
	}

}
