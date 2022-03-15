package d2m

import (
	"io/ioutil"
	"path/filepath"
)

func NewManifestFromDir(dir string) (manifest *Manifest, err error) {

	component, err := newComponentDir(dir)
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
