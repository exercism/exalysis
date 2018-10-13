package isogram

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/testhelper"
	"github.com/tehsphinx/exalysis/track/isogram/tpl"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/1", suggestion: tpl.Unicode, expected: true},
	{path: "./solutions/2", suggestion: tpl.Unicode, expected: true},
	{path: "./solutions/2", suggestion: tpl.RegexInFunc, expected: true},
	{path: "./solutions/2", suggestion: tpl.MustCompile, expected: true},
	{path: "./solutions/2", suggestion: tpl.JustReturn, expected: true},
	{path: "./solutions/2", suggestion: tpl.NonExistingMapValue, expected: true},
	{path: "./solutions/2", suggestion: tpl.IsLetter, expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "isogram")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion.ID()),
			fmt.Sprintf("test failed: %+v", test))
	}
}
