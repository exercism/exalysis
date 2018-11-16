package paraletterfreq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/testhelper"
	"github.com/tehsphinx/exalysis/track/paraletterfreq/tpl"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/1", suggestion: tpl.WaitGroup, expected: true},
	{path: "./solutions/1", suggestion: tpl.WaitGroupNotNeeded, expected: false},
	{path: "./solutions/1", suggestion: tpl.WaitGroupAddOne, expected: false},
	{path: "./solutions/1", suggestion: tpl.BufferSizeLen, expected: true},
	{path: "./solutions/2", suggestion: tpl.WaitGroup, expected: false},
	{path: "./solutions/2", suggestion: tpl.WaitGroupNotNeeded, expected: true},
	{path: "./solutions/2", suggestion: tpl.WaitGroupAddOne, expected: true},
	{path: "./solutions/2", suggestion: tpl.BufferSizeLen, expected: false},
	{path: "./solutions/3", suggestion: tpl.BufferSizeLen, expected: false},
	{path: "./solutions/4", suggestion: tpl.BufferSizeLen, expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "letter")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}
