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

	"github.com/tehsphinx/clipboard"
	"github.com/tehsphinx/exalysis"
)

var (
	minConfidence = flag.Float64("min_confidence", 0.8, "golint: minimum confidence of a problem to print it")
	watch         = flag.Bool("watch", false, "watch starts exalysis to watch the clipboard for `exercism download ...` commands")
)

func main() {
	flag.Parse()
	exalysis.LintMinConfidence = *minConfidence

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

	fmt.Println("waiting for exercism download links on the clipboard...")

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

		fmt.Println("<--------- Exalysis --------->")

		sugg, approval := exalysis.GetSuggestions(path)

		fmt.Println("\n" + sugg)
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
		return "", err
	}

	path := strings.TrimSpace(string(b))
	if _, err := os.Stat(path); err != nil {
		return "", err
	}
	return path, err
}
