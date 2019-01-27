package main

import (
	"bytes"
	"fmt"
	"io"
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

	rp, wp := io.Pipe()
	go func() {
		fmt.Fprintf(wp, "Some text to print!")
		wp.Close()
	}()
	bufp := make([]byte, 1)
	for {
		n, err := rp.Read(bufp)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Error: ", err)
			}
		}
		fmt.Print(n, string(bufp)+"-")
	}
}
