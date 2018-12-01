package exam

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
)

// GolangCILint runs GolangCI-Lint on provided go files and adds suggestions to the response.
// This runs a lot of linters. Goal is to cover a lot for now. Later individual linters can be run
// directly and disabled in GolangCI-Linter.
func GolangCILint(_ *astrav.Folder, r *extypes.Response, pkgName string) bool {
	res, state, err := golangCILint()
	if err != nil {
		fmt.Println(aurora.Gray("golangci-lint:\t"), aurora.Brown("SKIPPED"), aurora.Gray("\t(is GolangCI-Lint installed?)"))
		fmt.Println(err)
		return true
	}

	if state.Success() {
		fmt.Println(aurora.Gray("golangci-lint:\t"), aurora.Green("OK"))
		return true
	}

	fmt.Println(aurora.Gray("v:\t"), aurora.Red("FAIL"))
	fmt.Println(res)
	return false
}

func golangCILint() (string, *os.ProcessState, error) {
	cmd := exec.Command("golangci-lint", "run", "--disable", "govet,megacheck")

	b, err := cmd.CombinedOutput()
	if err != nil && !strings.HasPrefix(err.Error(), "exit status") {
		return "", nil, err
	}

	return string(b), cmd.ProcessState, nil
}
