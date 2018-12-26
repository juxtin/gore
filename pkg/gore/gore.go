package gore

import (
	"fmt"

	"github.com/juxtin/gore/pkg/files"
	"github.com/juxtin/gore/pkg/graph"
	rel "github.com/juxtin/gore/pkg/relationships"
)

func buildGraph(fs *files.FS, gopath string, root string) *graph.Graph {
	startFile := fs.CacheFile(root)
	graph := graph.NewGraph()
	initBranches := rel.GetImports(startFile)
	for _, to := range initBranches {
		downstream := buildGraph(fs, gopath, to)
		graph.Merge(downstream)
	}
	return graph
}

func BuildGraph(gopath string, root string) string {
	fs := files.NewFS(gopath)
	graph := buildGraph(fs, gopath, root)
	return graph.Graphviz()
}

func Smoke(gopath string, root string) {
	fs := files.NewFS(gopath)
	discovered := files.DiscoverFiles(fs, root)
	for i, df := range discovered {
		fmt.Println("File", i, "-", df.FullPath, " (", df.SrcPath, ")", len(df.Contents), "bytes")
	}
}
