package src

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
)

func ZipFile(zipFileName string, files []string) error {
	// create extension .zip if it doesn't specified yet
	if !strings.Contains(zipFileName, ".zip") {
		zipFileName = zipFileName + ".zip"
	}
	
	// create a new zip file
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		// open a file that will be added to archive
		fileToZip, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		// get file information
		fileInfo, err := fileToZip.Stat()
		if err != nil {
			return err
		}
		
		// create a new file header
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}

		// set the name of the file in the zip archive
		header.Name = file

		// create a new zip file entry
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// copy the file contents to the zip entry
		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Zip file %s created succesfully.\n", zipFileName)
	return nil
}
