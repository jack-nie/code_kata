package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(f)
	f.Close()
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fileInfo)
	err = os.Rename("test.txt", "test1.txt")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("test1.txt", os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file.Name())
	file.Close()
}
