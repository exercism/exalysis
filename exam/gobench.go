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

var (
	// Benchmarks setting whether to run the benchmarks or just the tests
	Benchmarks bool
)

// GoBench runs `go test` on provided path and adds suggestions to the response
func GoBench(_ *astrav.Folder, r *extypes.Response, _ string) bool {
	if !Benchmarks {
		fmt.Println(aurora.Gray("benchmarks:\t"), aurora.Brown("SKIPPED"))
		return true
	}

	res, state := bench()

	if state.Success() {
		fmt.Println(aurora.Gray("benchmarks:\t"), aurora.Green("OK"))
	} else {
		fmt.Println(aurora.Gray("benchmarks:\t"), aurora.Red("FAIL"))
	}

	fmt.Println(res)

	if state.Success() {
		return true
	}
	r.AppendTodo(gtpl.PassTests)
	return false
}

func bench() (string, *os.ProcessState) {
	cmd := exec.Command("go", "test", "-v", "--bench", ".", "--benchmem", "-test.run=^$")

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error running go test: ", err)
	}

	return string(b), cmd.ProcessState
}
