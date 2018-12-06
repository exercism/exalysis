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
	assert.Regexp(t, "Welcome to Exercism", string(output))
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
	assert.Regexp(t, "does not compile", string(output))
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
	assert.Regexp(t, "might find interesting", string(output))
}
