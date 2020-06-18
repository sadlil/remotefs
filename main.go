package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var progName = filepath.Base(os.Args[0])

func main() {
	log.SetFlags(0)
	log.SetPrefix(progName + ": ")

	flag.Parse()

	if flag.NArg() != 1 {
		os.Exit(2)
	}
	log.Println("Starting...")

	err := mount(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
}
