package src

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ZipFolder(sourceDir, destinationZip string) error {
	// create extension .zip if it doesn't specified yet
	if !strings.Contains(destinationZip, ".zip") {
		destinationZip = destinationZip + ".zip"
	}
	
	// Create a new zip file
	zipFile, err := os.Create(destinationZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// create a zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// walk throughh the source directory
	filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		// Handle errors during walking
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Ccreate a relative path for the current file or directory
		relativePath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// create a zip file header from file info
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// set the name of the file within the zip archive
		header.Name = relativePath

		// Create a new writer for the file in the zip archive
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// If it's a file, copy its contents to the zip archive
		if !info.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		}

		// continue walkingg
		return nil
	})

	fmt.Printf("Folder %v archived successfully: %v\n", sourceDir, destinationZip)
	return nil
}

