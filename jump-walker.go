package main

import (
	"flag"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/karrick/godirwalk"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/k0kubun/go-ansi"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var verbose = false
var terminalWidth = uint(80)

func PrintOverwrite(path string) {
	if verbose {
		fmt.Println("- ", path)
	} else {
		ansi.EraseInLine(2)
		ansi.CursorHorizontalAbsolute(0)
		len := uint(len(path))
		if (len > terminalWidth-2) {
			len = terminalWidth-2
		}
		sPath := path[0:len]
		fmt.Print("- ", sPath)
	}
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
	path = strings.Replace(path, "\\", "/", -1)
	ends := strings.Split(path, "/")[0]
	return path[0] == '.' || contains(postpone, ends)
}

func walk(cwd string, search string, checkLater []string) error {
	err := godirwalk.Walk(cwd, &godirwalk.Options{
		Callback: func(path string, info *godirwalk.Dirent) error {
			if info.IsDir() {
				PrintOverwrite(path)
				if info.Name() == search {
					if badName(path) {
						checkLater = append(checkLater, path)
						return filepath.SkipDir
					} else {
						elapsed := time.Since(start)
						fmt.Printf("\nFound [%s] in %s\n", path, elapsed)
						fmt.Println(path)
						os.Chdir(path)
						os.Exit(0)
					}
				}
			}
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			// Your program may want to log the error somehow.
			// fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)

			// For the purposes of this example, a simple SkipNode will suffice,
			// although in reality perhaps additional logic might be called for.
			return godirwalk.SkipNode
		},
		Unsorted: true,
	})

	if err != nil {
		//panic("Walking file system failed")
		//fmt.Println("Unable to check", cwd)
		return errors.Errorf("Unable to check " + cwd)
	}
	return err
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func checkByWalking(path string, search string, checkLater []string) error {
	fmt.Println("Checking", path)
	// store bad paths here which will be checked later
	err := walk(path, search, checkLater)
	fmt.Println("\ncheckLater:", len(checkLater))

	fmt.Println("[", search, "] is not found. Searching hidden folders...")
	for _, path := range checkLater {
		fmt.Println("Searching in", path)
		err = walk(path, search, checkLater)
	}
	return err
}

var start time.Time

func main() {
	start = time.Now()

	if len(os.Args) == 1 {
		panic("Usage: > jump some-folder")
	}
	var search = os.Args[1]
	if search == "" {
		panic("Usage: > jump [-v (verbose)] some-folder")
	}
	fmt.Printf("Searching %s...\n", search)

	var cwd, err = os.Getwd()
	if err != nil {
		panic("Unable to get cwd")
	}

	verbose = *flag.Bool("v", false, "verbose")
	fmt.Println("Verbose: ", verbose)

	terminalWidth, _ = terminal.Width()

	cwd = strings.Replace(cwd, "\\", "/", -1)
	var rootPath = strings.Split(cwd, "/")
	fmt.Println("Root Path: ", rootPath)
	var rootPathList []string
	for i, _ := range rootPath {
		var path = rootPath[0:i]
		rootPathList = append(rootPathList, strings.Join(path, "/"))
	}
	rootPathList = reverse(rootPathList)
	fmt.Println("rootPathList", rootPathList)

	var checkLater []string
	for _, path := range rootPathList {
		err = checkByWalking(path, search, checkLater)
		if err != nil {
			fmt.Println(err.(*errors.Error).ErrorStack())
		}
	}

	fmt.Println(`Not found, sorry ¯\_(ツ)_/¯`)
}
