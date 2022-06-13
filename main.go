package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Extracting html filepath
	indexPath, err := extractIndexPathFromArguments()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Processing html file
	err = process(indexPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func process(indexPath string) error {
	// Validate html filepath
	err := checkHtmlFile(indexPath)
	if err != nil {
		return err
	}

	// Reading html file content
	contentBytes, err := os.ReadFile(indexPath)
	if err != nil {
		return err
	}
	htmlContent := string(contentBytes)

	// Embedding all JavaScript into the html file
	htmlContent, err = embedJavaScriptFiles(htmlContent)
	if err != nil {
		return err
	}

	// Embedding all CSS into the html file
	htmlContent, err = embedCssFiles(htmlContent)
	if err != nil {
		return err
	}

	// Writing a Go file including all the html content
	err = writeHtmlContentToGoFile(htmlContent, indexPath)
	if err != nil {
		return err
	}

	return nil
}

func extractIndexPathFromArguments() (string, error) {
	if len(os.Args) != 2 {
		return "", errors.New("ERROR: Expected one command line argument: Path to the 'static' directory.")
	}
	return os.Args[1], nil
}

func checkHtmlFile(filePath string) error {
	if !strings.HasSuffix(filePath, ".html") {
		return fmt.Errorf("filepath does not end with '.html'")
	}
	if !fileExists(filePath) {
		return fmt.Errorf("filepath does not exist: %s", filePath)
	}
	return nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func embedJavaScriptFiles(htmlContent string) (string, error) {
	return htmlContent, nil
}

func embedCssFiles(htmlContent string) (string, error) {
	return htmlContent, nil
}

func writeHtmlContentToGoFile(htmlContent string, filePath string) error {
	return nil
}
