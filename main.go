package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.Flags())

	logger.Printf("We <3 Cement")
	logger.Printf("")
}
