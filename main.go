package main

import (
	"flag"
	"fmt"

	"github.com/davenicoll/terratidy/pkg"
)

func main() {
	var dirPath string
	flag.StringVar(&dirPath, "folder", "", "The folder path to apply DirectoryAutoFix")
	flag.Parse()

	if dirPath == "" {
		fmt.Println("Please provide a folder path using the -folder flag.")
		return
	}

	err := pkg.DirectoryAutoFix(dirPath)
	if err != nil {
		fmt.Println("Error during DirectoryAutoFix:", err)
		return
	}

	fmt.Println("DirectoryAutoFix completed successfully.")
}
