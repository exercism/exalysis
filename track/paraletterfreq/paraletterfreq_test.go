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
	{path: "./solutions/1", suggestion: tpl.RangeChan, expected: false},
	{path: "./solutions/1", suggestion: tpl.BufferSizeLen, expected: true},
	{path: "./solutions/1", suggestion: tpl.SelectNotNeeded, expected: true},
	{path: "./solutions/1", suggestion: tpl.GoroutineLeak, expected: true},
	{path: "./solutions/2", suggestion: tpl.WaitGroup, expected: false},
	{path: "./solutions/2", suggestion: tpl.WaitGroupNotNeeded, expected: true},
	{path: "./solutions/2", suggestion: tpl.WaitGroupAddOne, expected: true},
	{path: "./solutions/2", suggestion: tpl.RangeChan, expected: false},
	{path: "./solutions/2", suggestion: tpl.BufferSizeLen, expected: false},
	{path: "./solutions/2", suggestion: tpl.SelectNotNeeded, expected: true},
	{path: "./solutions/3", suggestion: tpl.BufferSizeLen, expected: false},
	{path: "./solutions/3", suggestion: tpl.ForRangeNoVars, expected: true},
	{path: "./solutions/4", suggestion: tpl.BufferSizeLen, expected: true},
	{path: "./solutions/5", suggestion: tpl.CombineMapsWhileWaiting, expected: true},
	{path: "./solutions/6", suggestion: tpl.ForRangeNoVars, expected: true},
	{path: "./solutions/7", suggestion: tpl.Mutex, expected: true},
	{path: "./solutions/7", suggestion: tpl.WaitGroupAddOne, expected: true},
	{path: "./solutions/8", suggestion: tpl.RangeChan, expected: true},
	{path: "./solutions/8", suggestion: tpl.WaitGroupAddOne, expected: true},
	{path: "./solutions/9", suggestion: tpl.ForRangeNoVars, expected: true},
	{path: "./solutions/10", suggestion: tpl.WaitGroup, expected: true},
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

/*
More solutions to look through:

https://exercism.io/mentor/solutions/01c8d3773ac941a1aba4ac40583c0b04
https://exercism.io/mentor/solutions/6be812e536484fe8b4cf81dadc14e532
https://exercism.io/mentor/solutions/8abc3969b5834267bf564e4ce421132a
https://exercism.io/mentor/solutions/916f4e9532e24367a1d39116e9a81d40
https://exercism.io/mentor/solutions/a90a88b7f4964251b1de8cb16c1cd03e
https://exercism.io/mentor/solutions/061b04cb2d6e4518b06971306ac7adbf
https://exercism.io/mentor/solutions/0980af90bb6649c7ab0453a7abde0fe0
https://exercism.io/mentor/solutions/70c5fb9aa5df4950915b30b721850136
https://exercism.io/mentor/solutions/5c68967dc4d245a0aee1c5694dd54c41
*/
