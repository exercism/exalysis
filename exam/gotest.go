package exam

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/logrusorgru/aurora"
	"github.com/tehsphinx/astrav"
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
func GoTest(_ *astrav.Folder, r *extypes.Response, pkgName string) (bool, error) {
	succTest, err := goTest(r)
	succTestRace := goTestRace(r, pkgName, err != nil)

	return succTest && succTestRace, err
}

func goTest(r *extypes.Response) (bool, error) {
	res, state, err := test()

	if state.Success() {
		fmt.Println(aurora.Gray("go test:\t"), aurora.Green("OK"))
		return true, nil
	}

	fmt.Println(aurora.Gray("go test:\t"), aurora.Red("FAIL"))
	fmt.Println(res)
	r.AppendTodoTpl(gtpl.PassTests)
	return false, err
}

func test() (string, *os.ProcessState, error) {
	cmd := exec.Command("go", "test")

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error running go test: ", err)
	}

	return string(b), cmd.ProcessState, err
}

func goTestRace(r *extypes.Response, pkgName string, skip bool) bool {
	if skipRace(pkgName) || skip {
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
	r.AppendTodoTpl(gtpl.RaceCondition)
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
