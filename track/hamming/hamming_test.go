package hamming

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
	{path: "./solutions/1", suggID: "error-message-format.md", expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "hamming")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggID), fmt.Sprintf("test failed: %+v", test))
	}
}
