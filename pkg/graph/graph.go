package graph

import "github.com/awalterschulze/gographviz"

type relationships map[string][]string

type Graph struct {
	relationships
}

func NewGraph() *Graph {
	r := make(map[string][]string)
	return &Graph{r}
}

func (g *Graph) AddImport(from string, to string) {
	existing, ok := g.relationships[from]
	if ok {
		g.relationships[from] = append(existing, to)
	} else {
		g.relationships[from] = []string{to}
	}
}

func (g *Graph) Graphviz() string {
	graphAst, _ := gographviz.ParseString(`digraph G {}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}
	for k, vs := range g.relationships {
		graph.AddNode("G", k, nil)
		for _, v := range vs {
			graph.AddNode("G", v, nil)
			graph.AddEdge(k, v, true, nil)
		}
	}
	return graph.String()
}
