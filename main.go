package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/basic/account-service/config"
	"github.com/basic/account-service/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	// db.Init()
	server.Init()
}
