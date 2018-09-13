package defs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/suggestion/scrabble"
)

func TestSuggesterInterface(t *testing.T) {
	assert.Implements(t, (*scrabble.Scrabble)(nil), new(Suggester))
}
