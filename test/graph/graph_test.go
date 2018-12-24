package graph

import (
	"testing"

	"github.com/juxtin/gore/pkg/graph"
)

func TestGraph(t *testing.T) {
	g := graph.NewGraph()
	g.AddImport("main", "util")
	g.AddImport("main", "net")
	g.AddImport("main", "cli")
	g.AddImport("cli", "parse")
	g.AddImport("parse", "util")
	g.AddImport("net", "tcp")
	g.AddImport("net", "util")
	g.Graphviz()
	// no assertion here for now, I'll figure something out
}
