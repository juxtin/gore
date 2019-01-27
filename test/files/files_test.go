package files

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/juxtin/gore/pkg/files"
	"github.com/stretchr/testify/assert"
)

func TestFiles(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	fs := files.NewFS(gopath)
	exFile := string(fs.CacheFile(gopath + "/src/" + "github.com/juxtin/gore/pkg/files/files.go"))
	fmt.Println("File:", exFile)
	assert.True(t, strings.Contains(exFile, "package files"), exFile)
}

func TestFilters(t *testing.T) {
	// Empty Filter
	empty := files.NewFilter()
	assert.True(t, empty.Accept("whatever"), "The empty filter should accept all paths.")

	// Comp with rejection
	rejectFooStar := empty.Comp(func(path string) files.Decision {
		if strings.HasPrefix(path, "foo") {
			return files.Reject
		}
		return files.Allow
	})
	assert.False(t, rejectFooStar.Accept("foo"), "rejectFooStar should reject 'foo'")
	assert.False(t, rejectFooStar.Accept("football"), "rejectFooStar should reject 'football'")
	assert.True(t, rejectFooStar.Accept("whatever"), "rejectFooStar should still accept 'whatever'")

	// Comp with acceptance
	acceptFoo := rejectFooStar.Comp(func(path string) files.Decision {
		if path == "foo" {
			return files.Accept
		}
		return files.Allow
	})
	assert.True(t, acceptFoo.Accept("foo"), "acceptFoo should accept 'foo'")
	assert.False(t, acceptFoo.Accept("foosball"), "acceptFoo should reject 'foosball'")
	assert.True(t, acceptFoo.Accept("whatever"), "acceptFoo should still accept 'whatever'")
}
