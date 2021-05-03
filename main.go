package main

import (
	"log"
	"os"

	"github.com/1gkx/goget/internal/cmd"
)

func main() {

	if len(os.Args) != 3 {
		log.Fatal("usage: download url filename")
		os.Exit(1)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
