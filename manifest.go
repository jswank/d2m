package d2m

import (
	"io/ioutil"
	"mime"
	"path/filepath"
)

type Manifest struct {
	Timestamp   string      `json:"timestamp"`
	Version     string      `json:"version"`
	Coordinates Coordinates `json:"coordinates"`
	Files       []File      `json:"files"`
}

type Coordinates struct {
	Group    string `json:"group"`
	Artifact string `json:"artifact"`
	Version  string `json:"version"`
}

type File struct {
	Hashes   map[string]string `json:"hashes"`
	Size     int64             `json:"size"`
	Filename string            `json:"filename"`
	MimeType string            `json:"mime_type"`
}

func NewManifest(group, artifact, version string) *Manifest {
	m := Manifest{}
	m.Coordinates = Coordinates{Group: group, Artifact: artifact, Version: version}
	return &m
}

func NewManifestFromDir(dir string) (manifest *Manifest, err error) {

	component, err := newComponent(dir)
	if err != nil {
		return
	}
	// fmt.Printf("ComponentDir: %+v\n", component)
	manifest = NewManifest(component.ParsedPom.GroupID, component.ParsedPom.ArtifactID, component.ParsedPom.Version)

	for _, f := range component.Artifacts {
		file := NewFile(f.Name(), f.Size())
		for _, h := range f.Hashes {
			// Ext() results in .ext so remove the .
			alg := filepath.Ext(h.Name())[1:]
			val, e := ioutil.ReadFile(filepath.Join(dir, h.Name()))
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

func (m *Manifest) AddFile(entry File) {
	m.Files = append(m.Files, entry)
	return
}

func NewFile(filename string, size int64) File {
	e := File{Filename: filename, Size: size, Hashes: map[string]string{}}
	e.MimeType = mimeType(filename)
	return e
}

func (e *File) AddHash(alg, value string) {
	e.Hashes[alg] = value
	return
}

func mimeType(filename string) string {
	// try to figure out content type
	ext := filepath.Ext(filename)
	mt := mime.TypeByExtension(ext)

	if mt == "" {
		if ext == ".pom" {
			mt = "text/xml"
		} else if ext == ".module" {
			mt = "application/json"
		} else {
			mt = "text/plain"
		}
	}

	return mt
}
