package main

import (
	"log"
	"os"
	"path"

	"github.com/leonhfr/search-go/pkg/gen"
)

const (
	searchPath = "wasm/main.go"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected one argument")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get current directory: %s", err)
	}

	entries, err := gen.Entries(path.Join(cwd, os.Args[1]))
	if err != nil {
		log.Fatalf("could not get entries: %s", err)
	}

	code, err := gen.SearchCode(entries)
	if err != nil {
		log.Fatalf("could not get source code: %s", err)
	}

	err = gen.WriteCode(code, path.Join(cwd, searchPath))
	if err != nil {
		log.Fatalf("could not get source code: %s", err)
	}
}
