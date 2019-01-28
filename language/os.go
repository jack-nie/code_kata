package main

import (
	"log"
	"os"
	"time"
	"fmt"
)

func main() {
	if err := os.Chmod("test.go", 0777); err != nil {
		log.Fatal(err)
	}

	mtime := time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2007, time.March, 2, 4,5,6,0, time.UTC)
	if err := os.Chtimes("test.go", atime, mtime); err != nil {
		log.Fatal(err)
	}

	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "USER":
			return "Gopher"
		}
		return ""
	}
	fmt.Println(os.Expand("Good ${DAY_PART}, $USER!", mapper))

	fmt.Println(os.ExpandEnv("$USER lives in ${HOME}"))
	fmt.Printf("%s lives in %s\n", os.Getenv("USER"), os.Getenv("HOME"))
	fmt.Printf("current uid is %d\n", os.Getuid())
	pwd, _ := os.Getwd()
	fmt.Printf("pwd is %s\n", pwd)
	hostname, _ := os.Hostname()
	fmt.Printf("hostname is %s\n", hostname)

	filename := "gogogogo"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("file dose not exist")
	}
}
