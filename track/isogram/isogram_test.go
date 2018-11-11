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
	{path: "./solutions/1", suggestion: tpl.IfContinue, expected: false},
	{path: "./solutions/2", suggestion: tpl.Unicode, expected: true},
	{path: "./solutions/2", suggestion: tpl.RegexInFunc, expected: true},
	{path: "./solutions/2", suggestion: tpl.MustCompile, expected: true},
	{path: "./solutions/2", suggestion: tpl.JustReturn, expected: true},
	{path: "./solutions/2", suggestion: tpl.NonExistingMapValue, expected: true},
	{path: "./solutions/2", suggestion: tpl.IsLetter, expected: true},
	{path: "./solutions/2", suggestion: tpl.IfContinue, expected: false},
	{path: "./solutions/3", suggestion: tpl.IsLetter, expected: true},
	{path: "./solutions/3", suggestion: tpl.IfContinue, expected: false},
	{path: "./solutions/4", suggestion: tpl.IfContinue, expected: true},
	{path: "./solutions/5", suggestion: tpl.NonExistingMapValue, expected: false},
	{path: "./solutions/5", suggestion: tpl.UniversalIsLetter, expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "isogram")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}
