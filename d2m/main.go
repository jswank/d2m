package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
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

	var manifest *d2m.Manifest
	var err error

	// is "directory" a url?  if so, try accessing it
	u, err := url.Parse(directory)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if u.Scheme == "" {
		manifest, err = d2m.NewManifestFromDir(directory)
	} else {
		manifest, err = d2m.NewManifestFromURL(directory)
	}
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	m, _ := json.MarshalIndent(manifest, "  ", "  ")
	fmt.Println(string(m))
}
