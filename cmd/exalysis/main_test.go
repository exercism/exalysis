package main

import (
	"log"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exePath = "./exalysis"

func build() error {
	out, err := exec.Command("go", "build", "-o", exePath, "main.go").CombinedOutput()
	if err != nil {
		log.Fatalf("failed to build: %v\n%v", err, string(out))
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
		log.Fatal(err)
	}
	assert.Regexp(t, "Welcome to Exercism", string(output))
}
