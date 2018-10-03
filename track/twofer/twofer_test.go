package twofer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/testhelper"
)

var suggestTests = []struct {
	path     string
	suggID   string
	expected bool
}{
	{path: "./solutions/1", suggID: "plus-used.md", expected: true},
	{path: "./solutions/1", suggID: "stub.md", expected: true},
	{path: "./solutions/1", suggID: "missing-comment.md", expected: false},
	{path: "./solutions/1", suggID: "comment-format.md", expected: false},
	{path: "./solutions/1", suggID: "wrong-comment-format.md", expected: false},
	{path: "./solutions/1", suggID: "minimal-conditional.md", expected: false},
	{path: "./solutions/2", suggID: "plus-used.md", expected: false},
	{path: "./solutions/2", suggID: "stub.md", expected: true},
	{path: "./solutions/2", suggID: "minimal-conditional.md", expected: true},
	{path: "./solutions/3", suggID: "use-string-placeholder.md", expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "twofer")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggID), fmt.Sprintf("test failed: %+v", test))
	}
}
