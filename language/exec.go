package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-al")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())

	output, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))

	path, err := exec.LookPath("date")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
}
