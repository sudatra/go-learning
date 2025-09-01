package scanner

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Scan(folder string) {
	fmt.Printf("Found Folders: \n\n")

	repositories := recursiveScanFolder(folder)
	filePath := getDotFilePath()

	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string
	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()

			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)

				continue
			}

			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}

			scanGitFolders(folders, path)
		}
	}

	return folders
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}