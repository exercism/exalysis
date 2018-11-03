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

//GoTest runs `go test` on provided path and adds suggestions to the response
func GoTest(_ *astrav.Folder, r *extypes.Response, _ string) bool {
	resBench, state := benchmark()

	if state.Success() {
		fmt.Println(aurora.Gray("go test:\t"), aurora.Green("OK"))
	} else {
		fmt.Println(aurora.Gray("go test:\t"), aurora.Red("FAIL"))
	}
	fmt.Println(resBench)

	if state.Success() {
		return true
	}
	r.AppendTodo(gtpl.PassTests)
	return false
}

func benchmark() (string, *os.ProcessState) {
	cmd := exec.Command("go", "test", "-v", "-bench", ".", "--benchmem")
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error running go test -bench: ", err)
	}

	return string(b), cmd.ProcessState
}
