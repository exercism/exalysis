package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/exercism/exalysis"
	"github.com/exercism/exalysis/exam"
	"github.com/tehsphinx/clipboard"
)

var (
	minConfidence = flag.Float64("min_confidence", 0.8, "golint: minimum confidence of a problem to print it")
	watch         = flag.Bool("watch", false, "starts Exalysis in watch mode, waiting for `exercism download ...` commands in the clipboard")
	outputAnswer  = flag.Bool("output", false, "outputs the answer to the student. Only applies to watch mode, where the answer is usually suppressed")
	bench         = flag.Bool("bench", false, "runs the benchmarks and outputs the result")
)

func main() {
	flag.Parse()
	exam.LintMinConfidence = *minConfidence
	exam.Benchmarks = *bench

	if *watch {
		watchClipboard()
		return
	}

	sugg, approval := exalysis.GetSuggestions(".")

	fmt.Println("\n" + sugg)
	fmt.Print("\n\n" + approval)
	if err := clipboard.WriteAll(sugg); err != nil {
		log.Println(err)
	}
}

func watchClipboard() {
	ch, cancel := clipboard.Watch(100 * time.Millisecond)
	defer cancel()

	fmt.Println("Watching for Exercism download links on the clipboard...")

	for clip := range ch {
		cmdText, ok := checkExercismDownload(clip)
		if !ok {
			continue
		}

		path, err := downloadExercise(cmdText)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("<-- Exalysis: %s -->\n", path)

		sugg, approval := exalysis.GetSuggestions(path)

		if *outputAnswer {
			fmt.Println("\n" + sugg)
		}
		fmt.Print("\n\n" + approval)
		if err := clipboard.WriteAll(sugg); err != nil {
			log.Println(err)
		}
	}
}

var downloadRegex = regexp.MustCompile("^exercism download --uuid=[[:xdigit:]]+$")

func checkExercismDownload(clip string) (string, bool) {
	if downloadRegex.MatchString(clip) {
		return clip, true
	}
	return "", false
}

func downloadExercise(cmdString string) (string, error) {
	parts := strings.Split(cmdString, " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	b, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error executing: %s: %s", strings.Join(parts, " "), err)
	}

	path := strings.TrimSpace(string(b))
	if _, err := os.Stat(path); err != nil {
		return "", err
	}
	return path, err
}
