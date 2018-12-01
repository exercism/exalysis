package luhn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/testhelper"
	"github.com/tehsphinx/exalysis/track/luhn/tpl"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/1", suggestion: tpl.RegexInFunc, expected: true},
	{path: "./solutions/1", suggestion: tpl.RegexToFast, expected: true},
	{path: "./solutions/2", suggestion: tpl.RegexInFunc, expected: false},
	{path: "./solutions/2", suggestion: tpl.RegexToFast, expected: true},
	{path: "./solutions/3", suggestion: tpl.RegexToFast, expected: false},
	{path: "./solutions/4", suggestion: tpl.OneLoop, expected: true},
	{path: "./solutions/4", suggestion: tpl.RegexToFast, expected: false},
	{path: "./solutions/5", suggestion: tpl.RegexInFunc, expected: false},
	{path: "./solutions/5", suggestion: tpl.RegexToFast, expected: true},
	{path: "./solutions/6", suggestion: tpl.OneLoop, expected: false},
	{path: "./solutions/6", suggestion: tpl.RegexToFast, expected: false},
	{path: "./solutions/7", suggestion: tpl.RegexToFast, expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "luhn")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}
