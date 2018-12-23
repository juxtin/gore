package gore

import (
	"testing"

	"github.com/juxtin/gore/pkg/gore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportsInFile(t *testing.T) {
	inputs := []string{"../pkg/main/main.go"}
	expected := []string{"fmt", "github.com/juxtin/gore/pkg/gore", "os"}
	results, err := gore.ImportsInFile(inputs)
	require.Nil(t, err, "Error reading the test input file!")
	assert.Equal(t, expected, results)
}
