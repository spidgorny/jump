package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		panic("Usage: > jump some-folder")
	}
	var search = os.Args[1]
	if search == "" {
		panic("Usage: > jump some-folder")
	}
	fmt.Printf("Searching %s...\n", search)

	var cwd, err = os.Getwd()
	if err != nil {
		panic("Unable to get cwd")
	}

	err = filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println('-', path)
			if info.Name() == search {
				os.Chdir(path)
				os.Exit(0)
			}
		}
		return nil
	})

	if err != nil {
		panic("Walking file system failed")
	}
}
