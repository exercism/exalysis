package scrabble

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
	{path: "./solutions/1", suggID: "unicode.md", expected: true},
	{path: "./solutions/3", suggID: "unicode.md", expected: true},
	{path: "./solutions/2", suggID: "unicode_loop.md", expected: true},
	{path: "./solutions/3", suggID: "unicode_loop.md", expected: false},
	{path: "./solutions/2", suggID: "move_map.md", expected: true},
	{path: "./solutions/3", suggID: "move_map.md", expected: false},
	{path: "./solutions/4", suggID: "move_map.md", expected: true},
	{path: "./solutions/2", suggID: "maprune.md", expected: true},
	{path: "./solutions/2", suggID: "try_switch.md", expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "scrabble")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggID), fmt.Sprintf("test failed: %+v", test))
	}
}
