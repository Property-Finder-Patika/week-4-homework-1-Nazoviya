package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

// visit every directory and print if they are visited.
func visit(path string, _ fs.DirEntry, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func main() {
	// waklkDir walks the file tree rooted at root, calling fn
	// for each file or directory in the tree, including root.
	err := filepath.WalkDir(".", visit)
	fmt.Printf("filepath.WalkDir() returned %v\n", err)
}
