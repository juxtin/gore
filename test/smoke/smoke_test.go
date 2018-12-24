package gore

import "testing"

import (
	"github.com/juxtin/gore/pkg/gore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportsInFile(t *testing.T) {
	inputs := []string{"smoke_test.go"}
	expected := []string{"testing", "github.com/juxtin/gore/pkg/gore", "github.com/stretchr/testify/assert", "github.com/stretchr/testify/require"}
	results, err := gore.ImportsInFile(inputs)
	require.Nil(t, err, "Error reading the test input file!")
	assert.Equal(t, expected, results)
}
