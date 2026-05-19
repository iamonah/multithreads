package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// searching for the name of a file in a directory and its subdirectories using goroutines and mutexes to
// handle concurrent access to the matches slice.
var matches []string
var mutex sync.Mutex
var wg sync.WaitGroup

func fileSearch(root string, fileName string) {
	defer wg.Done()
	entries, _ := os.ReadDir(root)
	for _, entry := range entries {
		if entry.IsDir() {
			wg.Add(1)
			go fileSearch(filepath.Join(root, entry.Name()), fileName)
		} else if strings.Contains(entry.Name(), fileName) {
			mutex.Lock()
			matches = append(matches, filepath.Join(root, entry.Name()))
			mutex.Unlock()
		}
	}
}

func main() {
	wg.Add(1)
	go fileSearch("/home/iamonah/Documents/service", "README.md")
	wg.Wait()
	for _, match := range matches {
		fmt.Printf("File found:  %v \n", match)
	}
}
