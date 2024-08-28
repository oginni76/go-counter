package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := flag.String("dir", ".", "Directory to count lines in")
	flag.Parse()

	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		fmt.Printf("Directory '%s' does not exist\n", *dir)
		os.Exit(1)
	}

	totalLines := 0
	err := traverseDirectory(*dir, &totalLines)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total lines in directory '%s': %d\n", *dir, totalLines)
}

func traverseDirectory(dir string, totalLines *int) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && !shouldIgnoreFile(path) {
			lines, err := countLines(path)
			if err != nil {
				return err
			}
			*totalLines += lines
			fmt.Printf("File: %s, Lines: %d\n", path, lines)
		}
		return nil
	})
}

func countLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

func shouldIgnoreFile(path string) bool {
	ignoredFiles := []string{".gitignore", "package.json", "go.mod", "go.sum", "bufio.scanner", "node_modules"}
	for _, ignoredFile := range ignoredFiles {
		if strings.HasSuffix(path, ignoredFile) {
			return true
		}
	}
	return false
}