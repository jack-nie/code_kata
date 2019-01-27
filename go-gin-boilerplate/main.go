package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jack-nie/code_kata/go-gin-boilerplate/config"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/models"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/server"
)

func main() {
	env := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(-1)
	}

	flag.Parse()
	config.Init(*env)
	models.InitDatabase()
	server.Init()
}
