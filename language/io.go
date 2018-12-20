package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Hello world! \n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	r1 := strings.NewReader("First reader!\n")
	r2 := strings.NewReader("Second reader!\n")
	buf := make([]byte, 8)
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}

	r3 := strings.NewReader("Second reader!\n")
	if _, err := io.CopyN(os.Stdout, r3, 5); err != nil {
		log.Fatal(err)
	}

	r4, w := io.Pipe()
	go func() {
		fmt.Fprint(w, "Some text to be read.\n")
		w.Close()
	}()

	buf2 := new(bytes.Buffer)
	buf2.ReadFrom(r4)
	fmt.Println(buf2.String())

	r5 := strings.NewReader("Hello world, gogogo!\n")
	buf3, err := ioutil.ReadAll(r5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", buf3)

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	content, err := ioutil.ReadFile("io.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", content)
}
