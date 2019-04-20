package twofer

import (
	"fmt"
	"testing"

	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/exercism/exalysis/testhelper"
	"github.com/exercism/exalysis/track/twofer/tpl"
	"github.com/stretchr/testify/assert"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/1", suggestion: tpl.PlusUsed, expected: true},
	{path: "./solutions/1", suggestion: tpl.Stub, expected: true},
	{path: "./solutions/1", suggestion: tpl.MissingComment, expected: false},
	{path: "./solutions/1", suggestion: tpl.CommentFormat, expected: false},
	{path: "./solutions/1", suggestion: tpl.WrongCommentFormat, expected: false},
	{path: "./solutions/1", suggestion: tpl.MinimalConditional, expected: false},
	{path: "./solutions/2", suggestion: tpl.PlusUsed, expected: false},
	{path: "./solutions/2", suggestion: tpl.Stub, expected: true},
	{path: "./solutions/2", suggestion: tpl.MinimalConditional, expected: true},
	{path: "./solutions/3", suggestion: tpl.UseStringPH, expected: true},
	{path: "./solutions/4", suggestion: tpl.MinimalConditional, expected: false},
	{path: "./solutions/5", suggestion: tpl.GeneralizeName, expected: true},
	{path: "./solutions/6", suggestion: tpl.UseStringPH, expected: true},
	{path: "./solutions/6", suggestion: tpl.MinimalConditional, expected: true},
	{path: "./solutions/7", suggestion: tpl.StringsJoin, expected: true},
	{path: "./solutions/8", suggestion: tpl.MissingComment, expected: true},
	{path: "./solutions/9", suggestion: tpl.MissingComment, expected: true},
	{path: "./solutions/10", suggestion: tpl.WrongCommentFormat, expected: false},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "twofer")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}
