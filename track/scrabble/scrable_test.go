package scrabble

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/testhelper"
	"github.com/tehsphinx/exalysis/track/scrabble/tpl"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/1", suggestion: tpl.Unicode, expected: true},
	{path: "./solutions/3", suggestion: tpl.Unicode, expected: true},
	{path: "./solutions/2", suggestion: tpl.UnicodeLoop, expected: true},
	{path: "./solutions/3", suggestion: tpl.UnicodeLoop, expected: false},
	{path: "./solutions/2", suggestion: tpl.MoveMap, expected: true},
	{path: "./solutions/3", suggestion: tpl.MoveMap, expected: false},
	{path: "./solutions/4", suggestion: tpl.MoveMap, expected: true},
	{path: "./solutions/2", suggestion: tpl.MapRune, expected: true},
	{path: "./solutions/2", suggestion: tpl.TrySwitch, expected: true},
	{path: "./solutions/5", suggestion: tpl.FlattenMap, expected: true},
	{path: "./solutions/5", suggestion: tpl.TrySwitch, expected: false},
	{path: "./solutions/5", suggestion: tpl.TypeConversion, expected: false},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "scrabble")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}
