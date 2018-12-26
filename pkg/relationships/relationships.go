package relationships

import (
	"fmt"
	"io/ioutil"
	"strconv"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
)

func ParseFile(document []byte) *sitter.Tree {
	parser := sitter.NewParser()
	parser.SetLanguage(golang.GetLanguage())
	return parser.Parse(document)
}

func nodeSource(node *sitter.Node, document []byte) string {
	start := node.StartByte()
	end := node.EndByte()
	rawSrc := string(document[start:end])
	src, err := strconv.Unquote(rawSrc)
	if err != nil {
		return rawSrc
	}
	return src
}

func hasChildren(node *sitter.Node) bool {
	return node.NamedChildCount() != 0
}

func concat(x []string, y []string) []string {
	ret := x
	for _, s := range y {
		ret = append(ret, s)
	}
	return ret
}

func getStringLiterals(node *sitter.Node, document []byte) []string {
	var ret []string
	children := node.NamedChildCount()
	for i := 0; i < int(children); i++ {
		n := node.NamedChild(i)
		if hasChildren(n) {
			ret = concat(ret, getStringLiterals(n, document))
		} else if n.Type() == "interpreted_string_literal" {
			ret = append(ret, nodeSource(n, document))
		}
	}
	return ret
}

func isImport(node *sitter.Node) bool {
	return node.Type() == "import_declaration"
}

func printChildren(node *sitter.Node, document []byte) {
	children := node.NamedChildCount()
	for i := 0; i < int(children); i++ {
		fmt.Println(node.NamedChild(i))
		fmt.Println(nodeSource(node, document))
		fmt.Println()
	}
}

func GetImports(document []byte) []string {
	ret := []string{}
	tree := ParseFile(document)
	n := tree.RootNode()
	children := n.NamedChildCount()
	for i := 0; i < int(children); i++ {
		node := n.NamedChild(i)
		if isImport(node) {
			ret = concat(ret, getStringLiterals(node, document))
		}
	}
	return ret
}

func ImportsInFile(filenames []string) ([]string, error) {
	var ret []string
	for _, filename := range filenames {
		document, err := ioutil.ReadFile(filename)
		if err != nil {
			return ret, err
		}
		imports := GetImports(document)
		ret = concat(ret, imports)
	}
	return ret, nil
}
