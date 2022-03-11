package d2m

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_NewManifest(t *testing.T) {
	m := NewManifest("com.github.120011676", "cipher", "0.0.7")
	fmt.Printf("manifest: %+v\n", m)
}

func Test_NewFile(t *testing.T) {
	e := NewFile("cipher-0.0.7.pom", 1808)
	fmt.Printf("file: %+v\n", e)
}

func Test_FullManifest(t *testing.T) {
	manifest := NewManifest("com.github.120011676", "cipher", "0.0.7")

	file := NewFile("cipher-0.0.7.pom", 1808)
	file.AddHash("md5", "e943c90d2f8f532eb50c44233c5bbcb5")

	manifest.AddFile(file)
	//	manifest.Entries = append(manifest.Entries, file)
	fmt.Printf("manifest: %+v\n", manifest)

	b, _ := json.MarshalIndent(manifest, "", " ")
	fmt.Println(string(b))
}

func Test_MimeType(t *testing.T) {
	mt := mimeType("foo.pom")
	fmt.Printf("ext .pom => %s\n", mt)
}
