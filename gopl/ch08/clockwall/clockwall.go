package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	tokyo := flag.String("Tokyo", "", "address of tokyo time server")
	london := flag.String("London", "", "address of London time server")
	flag.Parse()
	hosts := []string{*tokyo, *london}
	for _, host := range hosts {
		go func(s string) {
			conn, err := net.Dial("tcp", s)
			if err != nil {
				log.Fatal(err)
			}

			defer conn.Close()
			mustCopy(os.Stdout, conn)

		}(host)

	}
	for {
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
