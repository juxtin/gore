package gore

import (
	"fmt"

	"github.com/juxtin/gore/pkg/files"
	"github.com/juxtin/gore/pkg/graph"
	rel "github.com/juxtin/gore/pkg/relationships"
)

type SourceFile struct {
	PackageName string
	ImportPath  string
	SrcPath     string
	Imports     []string
	Contents    []byte
}

func AnalyzeFile(df *files.DiscoveredFile) SourceFile {
	contents := df.Contents
	packageName := rel.GetPackageName(contents)
	importPath := df.ImportPath
	srcPath := df.SrcPath
	imports := rel.GetImports(contents)
	return SourceFile{packageName, importPath, srcPath, imports, contents}
}

func AnalyzeFiles(dfs *[]files.DiscoveredFile) []SourceFile {
	ret := []SourceFile{}
	for _, df := range *dfs {
		ret = append(ret, AnalyzeFile(&df))
	}
	return ret
}

func buildGraph(sourceFiles *[]SourceFile) graph.Graph {
	graph := graph.NewGraph()
	for _, sf := range *sourceFiles {
		for _, imported := range sf.Imports {
			graph.AddImport(sf.ImportPath, imported)
		}
	}
	return *graph
}

func BuildGraph(gopath string, root string) string {
	fs := files.NewFS(gopath)
	discovered := files.DiscoverFiles(fs, root)
	analyzed := AnalyzeFiles(&discovered)
	graph := buildGraph(&analyzed)
	return graph.Graphviz()
}

func Smoke(gopath string, root string) {
	fs := files.NewFS(gopath)
	discovered := files.DiscoverFiles(fs, root)
	for i, df := range discovered {
		fmt.Println("File", i, "-", df.FullPath, " (", df.SrcPath, ")", len(df.Contents), "bytes")
	}
	analyzed := AnalyzeFiles(&discovered)
	for i, sf := range analyzed {
		fmt.Println("Package", i, "-", sf.ImportPath, "Imports:", sf.Imports, "Bytes:", len(sf.Contents))
	}
	graph := buildGraph(&analyzed)
	fmt.Println("Graph:")
	fmt.Println(graph.Graphviz())
}
