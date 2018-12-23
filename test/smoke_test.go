package gore

import (
	"testing"

	"github.com/juxtin/gore/pkg/gore"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	inputs := []string{"../pkg/main/main.go"}
	expected := []string{"fmt", "github.com/juxtin/gore/pkg/gore", "os"}
	results, err := gore.ImportsInFile(inputs)
	if err != nil {
		require.Fail(t, "Error during test!", err)
	}
	require.Equal(t, expected, results)
}
