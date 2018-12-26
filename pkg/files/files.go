package files

import (
	"io/ioutil"
	"strings"

	"github.com/juxtin/gore/pkg/debug"
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
	// TODO: actually find the dang file in here!
	todo := "Actually find the dang file!"
	if isFullyQualified(path) {
		return fs.gopath + "/src/" + path + "/" + todo
	}
	return path
}

func (fs *FS) CacheFile(path string) []byte {
	fullPath := fs.Resolve(path)
	debug.Print("Full path:", fullPath)
	if existing, ok := fs.Files[fullPath]; ok {
		return existing
	}
	document, err := ioutil.ReadFile(fullPath)
	if err != nil {
		fs.Files[fullPath] = nil
	} else {
		fs.Files[fullPath] = document
	}
	return fs.Files[fullPath]
}
