package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jswank/d2m"
)

var directory = "."

func init() {}

func main() {
	// read dir name as arg
	if len(os.Args) > 1 {
		directory = os.Args[1]
	}

	manifest, err := d2m.NewManifestFromDir(directory)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	m, _ := json.MarshalIndent(manifest, "  ", "  ")
	fmt.Println(string(m))
}
