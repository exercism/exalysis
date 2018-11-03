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
	//Benchmarks setting whether to run the benchmarks or just the tests
	Benchmarks bool
)

//GoTest runs `go test` on provided path and adds suggestions to the response
func GoTest(_ *astrav.Folder, r *extypes.Response, _ string) bool {
	res, state := test()

	if state.Success() {
		fmt.Println(aurora.Gray("go test:\t"), aurora.Green("OK"))
	} else {
		fmt.Println(aurora.Gray("go test:\t"), aurora.Red("FAIL"))
	}
	if Benchmarks || !state.Success() {
		fmt.Println(res)
	}

	if state.Success() {
		return true
	}
	r.AppendTodo(gtpl.PassTests)
	return false
}

func test() (string, *os.ProcessState) {
	var cmd *exec.Cmd
	if Benchmarks {
		cmd = exec.Command("go", "test", "-v", "-bench", ".", "--benchmem")
	} else {
		cmd = exec.Command("go", "test")
	}

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error running go test: ", err)
	}

	return string(b), cmd.ProcessState
}
