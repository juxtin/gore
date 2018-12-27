// This package tests the gore package

// more commentary to make it trickier for the parser to find the package
// declaration
package gore

import "testing"

import (
	"github.com/juxtin/gore/pkg/files"
	rel "github.com/juxtin/gore/pkg/relationships"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportsInFile(t *testing.T) {
	inputs := []string{"smoke_test.go"}
	expected := []string{"testing", "github.com/juxtin/gore/pkg/files", "github.com/juxtin/gore/pkg/relationships", "github.com/stretchr/testify/assert", "github.com/stretchr/testify/require"}
	results, err := rel.ImportsInFile(inputs)
	require.Nil(t, err, "Error reading the test input file!")
	assert.Equal(t, expected, results)
}

func TestGetPackageName(t *testing.T) {
	fs := files.NewFS("dummy gopath")
	document := fs.CacheFile("smoke_test.go")
	expected := "gore"
	result := rel.GetPackageName(document)
	assert.Equal(t, expected, result)
}
