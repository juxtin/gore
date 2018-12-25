package gore

import "testing"

import (
	rel "github.com/juxtin/gore/pkg/relationships"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportsInFile(t *testing.T) {
	inputs := []string{"smoke_test.go"}
	expected := []string{"testing", "github.com/juxtin/gore/pkg/relationships", "github.com/stretchr/testify/assert", "github.com/stretchr/testify/require"}
	results, err := rel.ImportsInFile(inputs)
	require.Nil(t, err, "Error reading the test input file!")
	assert.Equal(t, expected, results)
}
