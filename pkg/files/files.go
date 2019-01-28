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
	return &FS{
		gopath: gopath,
		Files:  map[string][]byte{},
	}
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
	return DiscoveredFile{
		SrcPath:    srcPath,
		FullPath:   fullPath,
		ImportPath: importPath,
		Contents:   contents,
	}
}

func isGoFile(path string) bool {
	return strings.HasSuffix(path, ".go") &&
		!strings.HasPrefix(path, ".") &&
		!strings.HasSuffix(path, "_test.go")
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

func WriteFile(path string, contents string) error {
	bytes := []byte(contents)
	err := ioutil.WriteFile(path, bytes, 0644)
	return err
}

type Decision int

const (
	// Allow delegates decision to the next function in the chain
	Allow Decision = iota
	// Accept short circuits, returning true
	Accept
	// Reject short circuits, returning false
	Reject
)

type Filter struct {
	Accept func(string) bool
}

func acceptAll(_ string) bool {
	return true
}

func NewFilter() Filter {
	return Filter{Accept: acceptAll}
}

func (f Filter) Comp(fn func(string) Decision) Filter {
	newFn := func(path string) bool {
		switch fn(path) {
		case Allow:
			return f.Accept(path)
		case Accept:
			return true
		case Reject:
			return false
		default:
			panic("Filtering function returned an invalid value!")
		}
	}
	return Filter{Accept: newFn}
}

// AcceptIf returns an Accept decision if the expr is true, otherwise Allow
func AcceptIf(expr bool) Decision {
	if expr {
		return Accept
	}
	return Allow
}

// RejectIf returns a Reject decision if the expr is true, otherwise Allow
func RejectIf(expr bool) Decision {
	if expr {
		return Reject
	}
	return Allow
}

// Require returns an Allow decision if the expr is true, otherwise Reject
func Require(expr bool) Decision {
	if expr {
		return Allow
	}
	return Reject
}
