package d2m

import (
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/2000Slash/gopom"
)

func NewManifestFromURL(url string) (manifest *Manifest, err error) {

	component, err := newComponentURL(url)
	if err != nil {
		return
	}
	// fmt.Printf("ComponentDir: %+v\n", component)
	manifest = NewManifest(component.ParsedPom.GroupID, component.ParsedPom.ArtifactID, component.ParsedPom.Version)

	for _, f := range component.Artifacts {
		file := NewFile(f.Name, f.Size)
		for _, h := range f.Hashes {
			// Ext() results in .ext so remove the .
			alg := filepath.Ext(h.Name)[1:]
			// fetch value from URL
			hashURL := *component.URL
			hashURL.Path = path.Join(component.URL.Path, h.Name)
			val, e := hashFromURL(hashURL.String())
			if e != nil {
				err = e
				return
			}
			file.AddHash(alg, string(val))
		}

		manifest.AddFile(file)
	}

	return
}

type componentURL struct {
	URL       *url.URL
	ParsedPom *gopom.Project
	Artifacts []artifactURL
}

type fileURL struct {
	Name string
	Size int64
}

type artifactURL struct {
	fileURL
	Hashes []fileURL
}

// read a maven-central directory index and create a conponentDir representing contents
func newComponentURL(u string) (comp componentURL, err error) {

	comp.URL, err = url.Parse(u)
	if err != nil {
		return
	}

	resp, err := http.Get(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		err = fmt.Errorf("non-2xx response: %d", resp.StatusCode)
		return
	}

	re, err := regexp.Compile(`<a href="(.+)"\s+title=".*">.*\s+(\d+)\s+$`)
	if err != nil {
		return
	}

	s := bufio.NewScanner(resp.Body)
	s.Split(bufio.ScanLines)
	files := []fileURL{}

	for s.Scan() {
		// if this looks like an dir index
		if re.MatchString(s.Text()) {
			matches := re.FindStringSubmatch(s.Text())
			file := fileURL{Name: matches[1]}
			size, _ := strconv.ParseInt(matches[2], 10, 64)
			file.Size = size
			files = append(files, file)
		}
	}

	hashes := []fileURL{}

	for _, f := range files {
		if filepath.Ext(f.Name) == ".pom" {
			// parse the POM file
			pomURL := *comp.URL
			pomURL.Path = path.Join(comp.URL.Path, f.Name)
			comp.ParsedPom, err = pomFromURL(pomURL.String())
			if err != nil {
				return
			}
		}
		if !isHash(filepath.Ext(f.Name)) {
			a := artifactURL{f, []fileURL{}}
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
		name := h.Name
		artifactName := name[0 : len(name)-len(filepath.Ext(name))]
		//		artifactName := h.Name()[0:len(h.Name() - len(filepath.Ext(h.Name())]
		// log.Printf("artifactName: %s", artifactName)
		for i, a := range comp.Artifacts {
			if artifactName == a.Name {
				// log.Printf("\tartifactName: %s, a.Name(): %s, h.Name(): %s\n", artifactName, a.Name(), h.Name())
				comp.Artifacts[i].Hashes = append(a.Hashes, h)
			}
		}
	}

	return

}

func pomFromURL(url string) (pom *gopom.Project, err error) {

	resp, err := http.Get(url)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		err = fmt.Errorf("non-2xx response: %d", resp.StatusCode)
		return
	}

	pom, err = pomReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func pomReadAll(r io.Reader) (*gopom.Project, error) {

	b, _ := ioutil.ReadAll(r)
	var project gopom.Project

	err := xml.Unmarshal(b, &project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func hashFromURL(url string) (hash string, err error) {

	resp, err := http.Get(url)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		err = fmt.Errorf("non-2xx response: %d", resp.StatusCode)
		return
	}

	hash, err = hashReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func hashReadAll(r io.Reader) (string, error) {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(b), nil

}
