package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	source := "D:\\Project\\Go\\source"
	target := "D:\\Project\\Go\\target"

	// Mengecek setiap file dan folder di direktori source & target
	sourceDir, err := os.Open(source)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sourceDir.Close()

	sourceFiles, err := sourceDir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	targetDir, err := os.Open(target)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer targetDir.Close()

	targetFiles, err := targetDir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range sourceFiles {
		filePathTarget := target + "/" + file.Name()
		filePathSource := strings.Replace(filePathTarget, target, source, 1)
		if _, err := os.Stat(filePathTarget); os.IsNotExist(err) {
			// bila hanya ada ditarget maka new
			fmt.Println(filePathSource + " NEW")
		}
	}

	for _, file := range targetFiles {
		filePathTarget := target + "/" + file.Name()
		filePathSource := strings.Replace(filePathTarget, target, source, 1)
		if _, err := os.Stat(filePathSource); os.IsNotExist(err) {
			// bila hanya ada di source maka delete
			fmt.Println(filePathTarget + " DELETED")
		}
	}
}
