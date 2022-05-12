package d2m

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/2000Slash/gopom"
)

type componentDir struct {
	ParsedPom *gopom.Project
	Timestamp time.Time
	Artifacts []artifact
}

type artifact struct {
	fs.FileInfo
	Hashes []fs.FileInfo
}

// read a directory and create a conponentDir representing contents
func newComponentDir(dir string) (comp componentDir, err error) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	var hashes []fs.FileInfo

	// set ParsedPom  Artifact
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".pom" {
			// parse the POM file
			comp.ParsedPom, err = gopom.Parse(filepath.Join(dir, f.Name()))
			if err != nil {
				return
			}
		}
		if !isHash(filepath.Ext(f.Name())) {
			a := artifact{f, []fs.FileInfo{}}
			comp.Artifacts = append(comp.Artifacts, a)
		} else {
			hashes = append(hashes, f)
		}
	}

	if comp.ParsedPom == nil {
		err = errors.New("No pom file found in the directory")
		return
	}

	// populate hashes
	for _, h := range hashes {
		// foo := "foo.asc"[0:3]
		name := h.Name()
		artifactName := name[0 : len(name)-len(filepath.Ext(name))]
		//		artifactName := h.Name()[0:len(h.Name() - len(filepath.Ext(h.Name())]
		// log.Printf("artifactName: %s", artifactName)
		for i, a := range comp.Artifacts {
			if artifactName == a.Name() {
				// log.Printf("\tartifactName: %s, a.Name(): %s, h.Name(): %s\n", artifactName, a.Name(), h.Name())
				comp.Artifacts[i].Hashes = append(a.Hashes, h)
			}
		}
	}

	return

}

// return true if the extension is a "hash"
func isHash(ext string) bool {
	hash_exts := []string{".md5", ".sha1", ".sha256", ".sha512", ".md5", ".asc"}
	for _, v := range hash_exts {
		if ext == v {
			return true
		}
	}
	return false
}
