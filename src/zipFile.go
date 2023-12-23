package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func ZipFile(zipFileName string, file *os.File) error {
	// create a new zip file
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// get file information
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	
	// create a new file header
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	// set the name of the file in the zip archive
	header.Name = fileInfo.Name()

	// create a new zip file entry
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// copy the file contents to the zip entry
	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}

	fmt.Printf("Zip file %s created succesfully.\n", zipFileName)
	return nil
}
