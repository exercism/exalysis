package scrabble

import (
	"github.com/tehsphinx/exalysis/extools"
	"golang.org/x/tools/go/loader"
)

//Scrabble build scrabble suggestions
func Scrabble(local *loader.PackageInfo, prog *loader.Program) []string {
	extools.PrintAST(local)

	var (
		sugg []string
	)

	sugg = testToLower(local, prog, sugg)
	sugg = testToUpper(local, prog, sugg)
	sugg = testMapRuneInt(local, prog, sugg)

	return sugg
}

const (
	speedComment = "- this exercise is a lot about speed and memory usage. If you haven't done so already " +
		"use `go test -v --bench . --benchmem` to benchmark your solution. Do that before and after every " +
		"to see if the change improved the speed or not. Also note the amount of memory allocations which " +
		"cost time. This exercise can be solved with 0 allocs/op."
	toLower = "- you could have a look at `unicode.ToLower` inside the for loop instead of `strings.ToUpper` " +
		"before the loop to increase speed"
	toUpper = "- you could have a look at `unicode.ToUpper` inside the for loop instead of `strings.ToUpper` " +
		"before the loop to increase speed"
	mapRune   = "- you could use a `map[rune]int` for direct lookup without type conversion instead of `map[string]int`"
	trySwitch = "- if you are up for it using a `switch` instead of a `map` will increase speed significantly"
)

var (
	speedCommentAdded bool
)

func testToLower(local *loader.PackageInfo, _ *loader.Program, sugg []string) []string {
	fn, ok := extools.GetUsageFunc("ToLower", local)
	if !ok {
		return sugg
	}
	if fn.Pkg().Name() == "unicode" {
		return sugg
	}

	sugg = addSpeedComment(sugg)
	return append(sugg, toLower)
}

func testToUpper(local *loader.PackageInfo, _ *loader.Program, sugg []string) []string {
	fn, ok := extools.GetUsageFunc("ToUpper", local)
	if !ok {
		return sugg
	}
	if fn.Pkg().Name() == "unicode" {
		return sugg
	}

	extools.GetDefinition("Score", local)
	//usage := extools.GetUsage("ToUpper", local)
	//dbg.Green(usage.Pos())

	sugg = addSpeedComment(sugg)
	return append(sugg, toUpper)
}

func testMapRuneInt(local *loader.PackageInfo, _ *loader.Program, sugg []string) []string {
	for _, t := range local.Types {
		switch t.Type.String() {
		case "map[rune]int":
			sugg = addSpeedComment(sugg)
			sugg = append(sugg, trySwitch)
			return sugg
		case "map[string]int":
			sugg = addSpeedComment(sugg)
			sugg = append(sugg, mapRune)
			sugg = append(sugg, trySwitch)
			return sugg
		}
	}
	return sugg
}

func addSpeedComment(sugg []string) []string {
	if !speedCommentAdded {
		speedCommentAdded = true
		return append(sugg, speedComment)
	}
	return sugg
}
