package main

import (
	"bufio"
	"os"
	"path/filepath"
)

// CreateDirectory creates directory if not exists.
func CreateDirectory() {
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		err := os.MkdirAll(Path, os.ModePerm)
		if err != nil {
			Fatalf("Error creating directory: %v", err)
		}
	}
}

// WriteToFile creates a file ( if not exists ), append the content and then close the file
func WriteToFile(filename string, message string) error {
	// open files r, w mode
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	// Close the file
	defer file.Close()

	// Append the message or content to be written
	if _, err = file.WriteString(message); err != nil {
		return err
	}

	return nil
}

// ListFile lists all the backup sql file to recreate the constraints
func ListFile(dir, suffix string) ([]string, error) {
	return filepath.Glob(filepath.Join(dir, suffix))
}

// ReadFile reads the file content and send it across
func ReadFile(filename string) ([]string, error) {
	var contentSaver []string

	// Open th file
	file, err := os.Open(filename)
	if err != nil {
		Fatalf("Error opening the file: %v", err)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contentSaver = append(contentSaver, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return contentSaver, err
	}
	return contentSaver, nil
}

// CurrentDir provides the current working directory
func CurrentDir() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		Fatalf("Error when trying to get the current directory, err: %v", err)
	}
	return cwd
}
