package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Printf("No dots: %q\n", filepath.Ext("index"))
	fmt.Printf("one dots: %q\n", filepath.Ext("index.js"))
	fmt.Printf("two dots: %q\n", filepath.Ext("index.js.je"))

	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"

	fmt.Println("On Unix:")
	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}
}
