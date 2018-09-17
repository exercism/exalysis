package scrabble

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/defs"
)

func TestSuggesterInterface(t *testing.T) {
	assert.Implements(t, (*Scrabble)(nil), new(defs.Suggester))
}
