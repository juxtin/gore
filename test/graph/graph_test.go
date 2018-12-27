package graph

import (
	"strings"
	"testing"

	"github.com/juxtin/gore/pkg/graph"
	"github.com/stretchr/testify/assert"
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
	s := g.Graphviz()
	// The order of lines returned by gographviz is non-deterministic, so we
	// need to test it by set comparison.
	results := newStringSet()
	for _, line := range strings.Split(s, "\n") {
		results.add(strings.TrimSpace(line))
	}
	expected := []string{
		"\"main\"->\"util\";",
		"\"main\"->\"net\";",
		"\"main\"->\"cli\";",
		"\"cli\"->\"parse\";",
		"\"net\"->\"tcp\";",
		"\"net\"->\"util\";",
		"\"cli\";",
		"\"main\";",
		"\"net\";",
		"\"parse\";",
		"\"tcp\";",
		"\"util\";",
	}
	for _, v := range expected {
		assertContains(t, results, v)
	}
}

// This language wisely lacks generics, which gives me the exciting opportunity
// to re-implement sets on an ad-hoc basis
type stringSet map[string]bool

func assertContains(t *testing.T, s stringSet, k string) {
	assert.Equal(t, s.contains(k), true, "Result did not contain '"+k+"'")
}

func newStringSet() stringSet {
	return make(map[string]bool)
}

func (s stringSet) add(k string) stringSet {
	s[k] = true
	return s
}

func (s stringSet) contains(key string) bool {
	_, ok := s[key]
	return ok
}
