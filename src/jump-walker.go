package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// this may need extending
// @see https://github.com/github/gitignore
func badName(path string) bool {
	postpone := []string{
		".svn",
		".hg",
		".git",
		"vendor",
		"node_modules",
		"__pycache__",
		".vagrant",
		"tmp",
		"temp",
	}
	ends := strings.Split(path, "/")[0]
	return path[0] == '.' || contains(postpone, ends)
}

func walk(cwd string, search string, checkLater []string) {
	err := filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Print("\r- ", path)
			if info.Name() == search {
				if badName(path) {
					checkLater = append(checkLater, path)
					return filepath.SkipDir
				} else {
					elapsed := time.Since(start)
					fmt.Printf("\nFound [%s] in %s", path, elapsed)
					fmt.Println(path)
					os.Chdir(path)
					os.Exit(0)
				}
			}
		}
		return nil
	})

	if err != nil {
		panic("Walking file system failed")
	}
}

var start time.Time

func main() {
	start = time.Now()

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

	// store bad paths here which will be checked later
	var checkLater []string
	walk(cwd, search, checkLater)

	fmt.Println("[", search, "] is not found. Searching hidden folders...")
	for _, path := range checkLater {
		fmt.Println("Searching in ", path)
		walk(path, search, checkLater)
	}

	fmt.Println(`Not found, sorry ¯\_(ツ)_/¯`)
}
