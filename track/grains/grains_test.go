package grains

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/testhelper"
	"github.com/tehsphinx/exalysis/track/grains/tpl"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/1", suggestion: tpl.BitsRotate64, expected: true},
	{path: "./solutions/1", suggestion: tpl.BitsRotate, expected: true},
	{path: "./solutions/1", suggestion: tpl.StaticTotalNice, expected: true},
	{path: "./solutions/1", suggestion: tpl.ErrorFormatted, expected: false},
	{path: "./solutions/2", suggestion: tpl.MathPow, expected: true},
	{path: "./solutions/2", suggestion: tpl.StaticTotal, expected: true},
	{path: "./solutions/2", suggestion: tpl.ErrorFormatted, expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "grains")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}
