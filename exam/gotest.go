package exam

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/logrusorgru/aurora"
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
)

var skipRaceExercises = []string{
	"twofer",
	"hamming",
	"raindrops",
	"scrabble",
	"isogram",
	"diffsquares",
	"luhn",
	"grains",
	"clock",
	"erratum",
}

// GoTest runs `go test` on provided path and adds suggestions to the response
func GoTest(_ *astrav.Folder, r *extypes.Response, pkgName string) bool {
	succTest := goTest(r)
	succTestRace := goTestRace(r, pkgName)

	return succTest && succTestRace
}

func goTest(r *extypes.Response) bool {
	res, state := test()

	if state.Success() {
		fmt.Println(aurora.Gray("go test:\t"), aurora.Green("OK"))
		return true
	}

	fmt.Println(aurora.Gray("go test:\t"), aurora.Red("FAIL"))
	fmt.Println(res)
	r.AppendTodo(gtpl.PassTests)
	return false
}

func test() (string, *os.ProcessState) {
	cmd := exec.Command("go", "test")

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error running go test: ", err)
	}

	return string(b), cmd.ProcessState
}

func goTestRace(r *extypes.Response, pkgName string) bool {
	if skipRace(pkgName) {
		fmt.Println(aurora.Gray("go test -race:\t"), aurora.Brown("SKIPPED"))
		return true
	}
	res, state := testRace()

	if state.Success() {
		fmt.Println(aurora.Gray("go test -race:\t"), aurora.Green("OK"))
		return true
	}

	fmt.Println(aurora.Gray("go test -race:\t"), aurora.Red("FAIL"))
	fmt.Println(res)
	r.AppendTodo(gtpl.RaceCondition)
	return false
}

func testRace() (string, *os.ProcessState) {
	cmd := exec.Command("go", "test", "-race")

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error running go test -race: ", err)
	}

	return string(b), cmd.ProcessState
}

func skipRace(pkgName string) bool {
	for _, n := range skipRaceExercises {
		if n == pkgName {
			return true
		}
	}
	return false
}
