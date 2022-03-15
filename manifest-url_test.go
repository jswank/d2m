package d2m

import "testing"

func Test_NewManifestFromURL(t *testing.T) {
	t.Log("test_NewManifestFromURL")
	url := "https://repo1.maven.org/maven2/com/github/120011676/cipher/0.0.7"

	manifest, err := NewManifestFromURL(url)
	if err != nil {
		t.Error(err)
	}

	t.Logf("manifest: %+v", manifest)
}

func Test_newComponentURL(t *testing.T) {
	t.Log("test_newComponentURL")
	url := "https://repo1.maven.org/maven2/com/github/120011676/cipher/0.0.7"
	comp, err := newComponentURL(url)
	if err != nil {
		t.Error(err)
	}

	t.Logf("componentURL: %+v", comp)
}

func Test_pomFromURL(t *testing.T) {
	t.Log("Test_pomFromURL")

	url := "https://repo1.maven.org/maven2/com/github/120011676/cipher/0.0.7/cipher-0.0.7.pom"

	pom, err := pomFromURL(url)
	if err != nil {
		t.Error(err)
	}

	if pom.GroupID != "com.github.120011676" {
		t.Errorf("unexpected GroupID: %s", pom.GroupID)
	}

	if pom.ArtifactID != "cipher" {
		t.Errorf("unexpected GroupID: %s", pom.ArtifactID)
	}

	if pom.Version != "0.0.7" {
		t.Errorf("unexpected Version: %s", pom.Version)
	}

	t.Logf("pom: %+v", pom)

}

func Test_hashFromURL(t *testing.T) {
	t.Log("Test_pomFromURL")

	url := "https://repo1.maven.org/maven2/com/github/120011676/cipher/0.0.7/cipher-0.0.7.pom.md5"

	hash, err := hashFromURL(url)
	if err != nil {
		t.Error(err)
	}

	if hash != "d6726cf0a3cf4db47c91f0786a04bae6" {
		t.Errorf("hash: %s", hash)
	}
}
