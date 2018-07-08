package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	port := flag.Int("port", 8080, "tcp port")
	flag.Parse()
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		if _, err := io.WriteString(conn, time.Now().Format("15:04:05\n")); err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
