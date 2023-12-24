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
	
	// Step 1: Create a new zip file
	zipFile, err := os.Create(destinationZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Step 2: Create a zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Step 3: Walk through the source directory
	filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		// Handle errors during walking
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Step 4: Create a relative path for the current file or directory
		relativePath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Step 5: Create a zip file header from file info
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Step 6: Set the name of the file within the zip archive
		header.Name = relativePath

		// Step 7: Create a new writer for the file in the zip archive
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Step 8: If it's a file, copy its contents to the zip archive
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

		// Step 9: Continue walking
		return nil
	})

	// Step 10: No need to return an error here, as we handle errors within the walk function
	fmt.Printf("Folder %v archived successfully: %v\n", sourceDir, destinationZip)
	return nil
}

