package scrabble

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/testhelper"
)

func TestUnicode(t *testing.T) {
	var solutions = []string{
		"./solutions/1",
	}
	for _, sol := range solutions {
		_, pkg, err := testhelper.LoadExample(sol, "scrabble")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		sugg := NewScrabble(pkg)
		sugg.Suggest(r)

		assert.True(t, r.HasSuggestion("unicode.md"))
	}
}

func TestUnicodeLoop(t *testing.T) {
	var solutions = []string{
		"./solutions/2",
	}
	for _, sol := range solutions {
		_, pkg, err := testhelper.LoadExample(sol, "scrabble")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		sugg := NewScrabble(pkg)
		sugg.Suggest(r)

		assert.True(t, r.HasSuggestion("unicode_loop.md"))
	}
}

func TestMoveMap(t *testing.T) {
	var solutions = []string{
		"./solutions/2",
	}
	for _, sol := range solutions {
		_, pkg, err := testhelper.LoadExample(sol, "scrabble")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		sugg := NewScrabble(pkg)
		sugg.Suggest(r)

		assert.True(t, r.HasSuggestion("move_map.md"))
	}
}

func TestMapRune(t *testing.T) {
	var solutions = []string{
		"./solutions/2",
	}
	for _, sol := range solutions {
		_, pkg, err := testhelper.LoadExample(sol, "scrabble")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		sugg := NewScrabble(pkg)
		sugg.Suggest(r)

		assert.True(t, r.HasSuggestion("maprune.md"))
	}
}

func TestTrySwitch(t *testing.T) {
	var solutions = []string{
		"./solutions/2",
	}
	for _, sol := range solutions {
		_, pkg, err := testhelper.LoadExample(sol, "scrabble")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		sugg := NewScrabble(pkg)
		sugg.Suggest(r)

		assert.True(t, r.HasSuggestion("try_switch.md"))
	}
}
