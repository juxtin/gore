package files

import (
	"io/ioutil"
	"strings"
)

type FS struct {
	gopath string
	Files  map[string][]byte
}

func NewFS(gopath string) *FS {
	return &FS{gopath, map[string][]byte{}}
}

func isFullyQualified(path string) bool {
	return strings.Contains(path, "/")
}

func (fs *FS) Resolve(path string) string {
	if isFullyQualified(path) {
		return fs.gopath + "/src/" + path
	}
	return path
}

func (fs *FS) CacheFile(path string) []byte {
	if existing, ok := fs.Files[path]; ok {
		return existing
	}
	document, err := ioutil.ReadFile(path)
	if err != nil {
		fs.Files[path] = nil
	} else {
		fs.Files[path] = document
	}
	return fs.Files[path]
}
