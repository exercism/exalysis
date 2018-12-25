package main

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exePath = "./exalysis"

func build() error {
	out, err := exec.Command("go", "build", "-o", exePath, "main.go").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to build: %s\n%s", err, out)
	}
	return nil
}

func TestHappyPath(t *testing.T) {
	if err := build(); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("../../" + exePath)
	cmd.Dir = "./testdata/happypath"
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("%s: %s", err, output)
	}
	assert.Contains(t, string(output), "Welcome to Exercism")
}

func TestCompileError(t *testing.T) {
	if err := build(); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("../../" + exePath)
	cmd.Dir = "./testdata/compile_error"
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("%s: %s", err, output)
	}
	assert.Contains(t, string(output), "does not compile")
}

func TestCompileError2(t *testing.T) {
	if err := build(); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("../../" + exePath)
	cmd.Dir = "./testdata/compile_error2"
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("%s: %s", err, output)
	}
	assert.Contains(t, string(output), "does not pass the tests")
	assert.NotContains(t, string(output), "race conditions")
	assert.NotContains(t, string(output), "`go vet`")
}

func TestVetError(t *testing.T) {
	if err := build(); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("../../" + exePath)
	cmd.Dir = "./testdata/vet_error"
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("%s: %s", err, output)
	}
	assert.NotContains(t, string(output), "`golint`")
	assert.Contains(t, string(output), "`go vet`")
}

func TestTip(t *testing.T) {
	if err := build(); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("../../" + exePath)
	cmd.Dir = "./testdata/happypath"
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("%s: %s", err, output)
	}
	assert.Contains(t, string(output), "might find interesting")
}
