package main

import (
	"log"

	"github.com/1gkx/goget/internal/arguments"
	"github.com/1gkx/goget/internal/cmd"
)

func main() {
	if err := cmd.Start(arguments.Parse()); err != nil {
		log.Fatal(err)
	}
}
