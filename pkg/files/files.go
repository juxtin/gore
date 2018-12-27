package files

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/juxtin/gore/pkg/debug"
)

type FS struct {
	gopath string
	Files  map[string][]byte
}

type DiscoveredFile struct {
	SrcPath    string
	FullPath   string
	ImportPath string
	Contents   []byte
}

func NewFS(gopath string) *FS {
	return &FS{gopath, map[string][]byte{}}
}

func importPath(srcPath string) string {
	dir, _ := path.Split(srcPath)
	return strings.TrimSuffix(dir, "/")
}

func discoverFile(fs *FS, fullPath string, info os.FileInfo) DiscoveredFile {
	contents := fs.CacheFile(fullPath)
	srcPath := strings.TrimPrefix(fullPath, fs.gopath+"/src/")
	importPath := importPath(srcPath)
	debug.Print("Shortened", fullPath, "to", importPath)
	return DiscoveredFile{srcPath, fullPath, importPath, contents}
}

func isGoFile(path string) bool {
	return strings.HasSuffix(path, ".go") && !strings.HasPrefix(path, ".") && !strings.HasSuffix(path, "_test.go")
}

func DiscoverFiles(fs *FS, rootDir string) []DiscoveredFile {
	ret := []DiscoveredFile{}
	walker := func(path string, info os.FileInfo, err error) error {
		if isGoFile(path) && err == nil {
			debug.Print("Discovered:", path)
			absPath, _ := filepath.Abs(path)
			ret = append(ret, discoverFile(fs, absPath, info))
		}
		return nil
	}
	filepath.Walk(rootDir, walker)
	return ret
}

func isFullyQualified(path string) bool {
	return strings.Contains(path, "/")
}

func (fs *FS) CacheFile(path string) []byte {
	debug.Print("Reading file at path:", path)
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
