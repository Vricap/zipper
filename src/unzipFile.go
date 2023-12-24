package src

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func UnzipFIle(source, destination string) error {
	// open the zip file (source)
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()
	
	// make a directory (destination) and grant user all access to the dir
	errMkDir := os.MkdirAll(destination, os.ModePerm)
	if errMkDir != nil {
		return errMkDir
	}

	// iterate each file in the arcive
	for _, file := range reader.File {
		// open each archived file for reading
		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		// construct the path where extracted file will be
		extractFilePath := filepath.Join(destination, file.Name)

		// check the file info
		if file.FileInfo().IsDir() {
			// if it's a directory, then make a directory
			os.MkdirAll(extractFilePath, os.ModePerm)
		} else {
			// if it's not a dir--a file, then create a file
			extractedFile, err := os.Create(extractFilePath)
			if err != nil {
				return err
			}
			defer extractedFile.Close()

			// copy the file CONTENT from the archived file to the newly created file
			_, errCopy := io.Copy(extractedFile, zippedFile)
			if errCopy != nil {
				return errCopy
			}
		}
	}

	fmt.Printf("File %v unzipped successfully: %v\n", source, destination)
	return nil
}