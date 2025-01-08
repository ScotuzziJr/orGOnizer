package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// File Organizer
//
// 1. Accept the directory path from the user
// 2. Read files
// 3. Categorize files by extension
// 4. Move files to category directories

func main() {
	// 1. Accept the directory path from the user
	if len(os.Args) < 2 {
		fmt.Println("Usage: OrGOnizer [directotry path]")
		return
	}

	dirPath := os.Args[1]

	fmt.Println("Organizing files in: ", dirPath)

	// 2. Read files
	files, err := readFiles(dirPath)

	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// 3. Categorize files by extension
	for _, file := range files {
		if !file.IsDir() {
			category := getCategory(file.Name())
			categoryPath := filepath.Join(dirPath, category)

			if err := createDir(dirPath, category); err != nil {
				fmt.Println("Error creating folder: ", err)
				continue
			}

			// 4. Move files to category directories
			srcPath := filepath.Join(dirPath, file.Name())

			if err := moveFile(srcPath, categoryPath); err != nil {
				fmt.Println("Error moving the file: ", err)
			} else {
				fmt.Printf("Moved: %s -> %s\n", file.Name(), category)
			}
		}
	}
}

func readFiles(dirPath string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dirPath)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func getCategory(fileName string) string {
	ext := filepath.Ext(fileName)

	switch ext {
	case ".jpg", ".jpeg", ".png", ".avif", ".wepb", ".gif", ".svg":
		return "Images"
	case ".mp4", ".mov", ".mkv":
		return "Videos"
	case ".psd", ".ai", ".xcf":
		return "Design"
	case ".txt", ".csv", ".docx", ".doc", ".xlsx", ".xls", ".pdf", ".odt", ".ods":
		return "Documents"
	default:
		return "Others"
	}
}

func createDir(dirPath, category string) error {
	categoryPath := filepath.Join(dirPath, category)

	return os.MkdirAll(categoryPath, os.ModePerm)
}

func moveFile(filePath, destDir string) error {
	destPath := filepath.Join(destDir, filepath.Base(filePath))

	return os.Rename(filePath, destPath)
}
