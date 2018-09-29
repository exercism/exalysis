package scrabble

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/testhelper"
)

func TestUnicode(t *testing.T) {
	_, pkg, err := testhelper.LoadExample("./solutions/1", "scrabble")
	if err != nil {
		t.Fatal(err)
	}

	r := extypes.NewResponse()
	sugg := NewScrabble(pkg)
	sugg.Suggest(r)

	assert.True(t, r.HasSuggestion("unicode.md"))
}
