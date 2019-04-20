package hamming

import (
	"fmt"
	"testing"

	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/exercism/exalysis/testhelper"
	"github.com/exercism/exalysis/track/hamming/tpl"
	"github.com/stretchr/testify/assert"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/1", suggestion: tpl.ErrorMessageFormat, expected: true},
	{path: "./solutions/1", suggestion: tpl.RuneToByte, expected: true},
	{path: "./solutions/1", suggestion: tpl.ReturnZeroValue, expected: true},
	{path: "./solutions/1", suggestion: tpl.DeclareNeeded, expected: false},
	{path: "./solutions/2", suggestion: tpl.MultiStringConv, expected: true},
	{path: "./solutions/2", suggestion: tpl.ReturnZeroValue, expected: false},
	{path: "./solutions/2", suggestion: tpl.DeclareNeeded, expected: true},
	{path: "./solutions/3", suggestion: tpl.ErrorMessageFormat, expected: true},
	{path: "./solutions/3", suggestion: tpl.Increase, expected: true},
	{path: "./solutions/3", suggestion: tpl.ReturnZeroValue, expected: true},
	{path: "./solutions/3", suggestion: tpl.DeclareNeeded, expected: false},
	{path: "./solutions/4", suggestion: tpl.InvertIf, expected: true},
	{path: "./solutions/4", suggestion: tpl.ReturnZeroValue, expected: true},
	{path: "./solutions/4", suggestion: tpl.DeclareNeeded, expected: false},
	{path: "./solutions/5", suggestion: tpl.InvertIf, expected: false},
	{path: "./solutions/5", suggestion: tpl.ReturnZeroValue, expected: true},
	{path: "./solutions/5", suggestion: tpl.DeclareNeeded, expected: true},
	{path: "./solutions/6", suggestion: tpl.DeclareNeeded, expected: true},
	{path: "./solutions/6", suggestion: tpl.ReturnZeroValue, expected: true},
	{path: "./solutions/7", suggestion: tpl.NoErrMsg, expected: true},
	{path: "./solutions/7", suggestion: tpl.ReturnZeroValue, expected: true},
	{path: "./solutions/8", suggestion: tpl.ReturnZeroValue, expected: true},
	{path: "./solutions/9", suggestion: tpl.NakedReturns, expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "hamming")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}
