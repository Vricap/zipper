package main

import (
	"os"
)

func main() {
	// create a what-to-be name of the zip archive
	zipFileName := "example.zip"

	// open a file to be archived
	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = ZipFile(zipFileName, file)
	if err != nil {
		panic(err)
	}
}
