package d2m

import (
	"mime"
	"path/filepath"
	"time"
)

const manifest_version = 1

type Manifest struct {
	Timestamp   string      `json:"timestamp"`
	Version     int64       `json:"version"`
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
	m := Manifest{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   manifest_version,
	}
	m.Coordinates = Coordinates{Group: group, Artifact: artifact, Version: version}
	return &m
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
