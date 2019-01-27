package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	addr, err := net.ResolveTCPAddr("tcp", "baidu.com:80")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Content: ", string(content))

	fmt.Println("status: ", status)
	conn.Close()
}
