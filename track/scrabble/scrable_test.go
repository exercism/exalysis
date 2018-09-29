package scrabble

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
)

func TestSuggesterInterface(t *testing.T) {
	assert.Implements(t, (*Scrabble)(nil), new(extypes.Suggester))
}
