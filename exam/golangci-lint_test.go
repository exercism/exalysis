package exam

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/exercism/exalysis/testhelper"
)

var golangCILintTests = []struct {
	path       string
	expected   bool
	suggestion gtpl.Template
	expectSugg bool
	pkgName    string
}{
	{path: "./solutions/0", expected: false},
	{path: "./solutions/1", expected: true, pkgName: "twofer"},
	{path: "./solutions/2", expected: true, pkgName: "hamming"},
	{path: "./solutions/3", expected: true},
	{path: "./solutions/4", expected: false},
}

func TestGolangCILint(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if r := os.Chdir(dir); r != nil {
			err = r
		}
	}()

	for _, test := range golangCILintTests {
		folder, _, err := testhelper.LoadFolder(path.Join(dir, test.path))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.Chdir(folder.GetPath()); err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		ok := GolangCILint(folder, r, test.pkgName)

		failMsg := fmt.Sprintf("test failed: %+v", test)
		assert.Equal(t, test.expected, ok, failMsg)
		if test.suggestion != nil {
			assert.Equal(t, test.expectSugg, r.HasSuggestion(test.suggestion), failMsg)
		}
	}
}
