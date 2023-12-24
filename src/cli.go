package src

import (
	"fmt"
	"os"
	"strings"
)

func Execute()  {
	// what is flag anyway
	if len(os.Args) < 2 {
		fmt.Println(`Please specify an action flag.
'-z' for zipping
'-u' for unzipping`)
		return
	}

	if os.Args[1] == "-z" {
		if len(os.Args) < 3 {
			fmt.Println(`Please specify type flag.
Available flag type:
'-fl' for file
'-fd' for folder`)
			return
		}

		// this is mad. i might refactor this dw
		switch os.Args[2] {
			case "-fl":
				if len(os.Args) == 3 {
					fmt.Println("Please specify the source path")
					return
				}
				// ZIPPING A FILE
				// slice of files name to be archived
				files := make([]string, 0)

				// create a what-to-be name of the zip archive
				zipFileName := "zipper.zip"

				if string(os.Args[3][0]) == "[" {
					i := 0
					str := ""
					for i = 3; string(os.Args[i][len(os.Args[i]) - 1]) != "]"; i++ { // this is retarded
						str = strings.ReplaceAll(strings.ReplaceAll(os.Args[i], "[", ""), "]", "")
						files = append(files, str)
						str = ""
					}
					str = strings.ReplaceAll(strings.ReplaceAll(os.Args[i], "[", ""), "]", "")
					files = append(files, str)
					if os.Args[i] != os.Args[len(os.Args) - 1] {
						zipFileName = os.Args[len(os.Args) - 1]
					}
				} else {
					files = append(files, os.Args[3])
				}
				if len(os.Args) == 5 && !strings.Contains(os.Args[len(os.Args) - 1], "]"){
					zipFileName = os.Args[len(os.Args) - 1]
				}
				doZipFile(zipFileName, files)

			case "-fd":
				if len(os.Args) == 3 {
					fmt.Println("Please specify the source path")
					return
				}
				// ZIPPING A FOLDER
				// set the path of folder to be zipped
				sourceFolder := os.Args[3]
				
				// // set destination path where zipped folder will reside
				destinationFolder := "zipper.zip"
				if os.Args[3] != os.Args[len(os.Args) - 1] {
					destinationFolder = os.Args[len(os.Args) - 1]
				}
				doZipFolder(sourceFolder, destinationFolder)
			
			default: 
			fmt.Println(`Your specified flag type is unavailable.
Available flag type:
'-fl' for file
'-fd' for folder`)
		}

		return
	}

	if os.Args[1] == "-u" {
		if len(os.Args) == 2 {
			fmt.Println("Please specify the source path")
			return
		}
		// UNZIPPING
		// set the path of file to be unzipped
		source := os.Args[2]

		// set destination path where unzipped file will reside
		destination := "zipper"
		if os.Args[2] != os.Args[len(os.Args) - 1] {
			destination = os.Args[len(os.Args) - 1]
		}
		doUnzip(source, destination)

		return
	}
}

// ZIPPING A FILE
// create a what-to-be name of the zip archive
// zipFileName := "/example/example.zip"

// slice of files name to be archived
// files := []string{"test.txt", "test1.txt"}
// doZipFile(zipFileName, files)


// ZIPPING A FOLDER
// set the path of folder to be zipped
// sourceFolder := "test"

// set destination path where zipped folder will reside
// destinationFolder := "done"
// doZipFolder(sourceFolder, destinationFolder)


// UNZIPPING
// set the path of file to be unzipped
// source := "ezyzip.zip"

// set destination path where unzipped file will reside
// destination := "."
// doUnzip(source, destination)


func doZipFile(zipFileName string, files []string)  {
	errZip := ZipFile(zipFileName, files)
	if errZip != nil {
	 	os.Remove(zipFileName)
		panic(errZip)
	}	
}

func doZipFolder(sourceFolder, destinationFolder string)  {
	errZipFd := ZipFolder(sourceFolder, destinationFolder)
	if errZipFd != nil {
		os.Remove(destinationFolder)
		panic(errZipFd)
	}	
}

func doUnzip(source, destination string) {
	errUnzip := UnzipFIle(source, destination)
	if errUnzip != nil {
		os.Remove(destination)		
		panic(errUnzip)
	}
}