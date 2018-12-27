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
