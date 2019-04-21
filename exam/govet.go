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

// GoVet runs go vet on provided go files and adds suggestions to the response
func GoVet(_ *astrav.Folder, r *extypes.Response, pkgName string, skip bool) bool {
	if skip {
		fmt.Println(aurora.Gray("go vet:\t\t"), aurora.Brown("SKIPPED"))
		return true
	}

	res, state := goVet()

	if state.Success() {
		fmt.Println(aurora.Gray("go vet:\t\t"), aurora.Green("OK"))
		return true
	}

	fmt.Println(aurora.Gray("go vet:\t\t"), aurora.Red("FAIL"))
	fmt.Println(res)

	if pkgName == "twofer" || pkgName == "hamming" {
		r.AppendImprovementTpl(gtpl.NotVetted)
	} else {
		r.AppendTodoTpl(gtpl.NotVetted)
	}

	return false
}

func goVet() (string, *os.ProcessState) {
	cmd := exec.Command("go", "vet")

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error running go vet: ", err)
	}

	return string(b), cmd.ProcessState
}
