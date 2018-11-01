package raindrops

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/testhelper"
	"github.com/tehsphinx/exalysis/track/raindrops/tpl"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/1", suggestion: tpl.ManyLoops, expected: true},
	{path: "./solutions/1", suggestion: tpl.ConcatNotNeeded, expected: false},
	{path: "./solutions/2", suggestion: tpl.ConcatNotNeeded, expected: true},
	{path: "./solutions/3", suggestion: tpl.StringsBuilder, expected: true},
	{path: "./solutions/3", suggestion: tpl.Itoa, expected: false},
	{path: "./solutions/4", suggestion: tpl.PlusEqual, expected: true},
	{path: "./solutions/5", suggestion: tpl.Itoa, expected: true},
	{path: "./solutions/6", suggestion: tpl.ExtensiveFor, expected: true},
	{path: "./solutions/6", suggestion: tpl.PlusEqual, expected: true},
	{path: "./solutions/7", suggestion: tpl.LoopMap, expected: true},
	{path: "./solutions/8", suggestion: tpl.ExtensiveFor, expected: true},
	{path: "./solutions/9", suggestion: tpl.ExtensiveFor, expected: true},
	{path: "./solutions/10", suggestion: tpl.ExtensiveFor, expected: true},
	{path: "./solutions/11", suggestion: tpl.ExtensiveFor, expected: false},
	{path: "./solutions/12", suggestion: tpl.StringsBuilder, expected: true},
	{path: "./solutions/13", suggestion: tpl.RemoveExtraBool, expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "raindrops")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}
