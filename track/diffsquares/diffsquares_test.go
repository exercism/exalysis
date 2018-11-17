package diffsquares

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/testhelper"
	"github.com/tehsphinx/exalysis/track/diffsquares/tpl"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/1", suggestion: tpl.SquareSumLoop, expected: true},
	{path: "./solutions/1", suggestion: tpl.SumSquareLoop, expected: true},
	{path: "./solutions/1", suggestion: tpl.CalcRangeCondition, expected: true},
	{path: "./solutions/2", suggestion: tpl.SquareSumLoop, expected: true},
	{path: "./solutions/2", suggestion: tpl.SumSquareLoop, expected: true},
	{path: "./solutions/2", suggestion: tpl.CalcRangeCondition, expected: false},
	{path: "./solutions/3", suggestion: tpl.MathPow, expected: true},
	{path: "./solutions/4", suggestion: tpl.SumSquareLoop, expected: true},
	{path: "./solutions/4", suggestion: tpl.SquareSumLoop, expected: true},
	{path: "./solutions/4", suggestion: tpl.MathPow, expected: true},
	{path: "./solutions/5", suggestion: tpl.Dry, expected: true},
	{path: "./solutions/6", suggestion: tpl.SquareSumLoop, expected: true},
	{path: "./solutions/6", suggestion: tpl.SumSquareLoop, expected: true},
	{path: "./solutions/6", suggestion: tpl.MathPow, expected: true},
	{path: "./solutions/6", suggestion: tpl.BasicFloat64, expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "diffsquares")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}
