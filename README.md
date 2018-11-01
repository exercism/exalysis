# Exalysis

[![Build Status](https://travis-ci.org/tehsphinx/exalysis.svg?branch=master)](https://travis-ci.org/tehsphinx/exalysis)

Exalysis is a tool ment for mentors of the [exercism](https://exercism.io) go track.

## Status
This is a prove of concept. I use it to mentor and enhance/extend it as I go.

***Exercise implementations:***
- TwoFer
- Hamming
- Raindrops
- Scrabble Score 

## Installation

```
go install github.com/tehsphinx/exalysis/cmd/exalysis
```

## Usage

### Watch mode

The watch mode you can start from anywhere:

```
~/ > exalysis --watch
```

or execute the source code directly:

```
path to cmd/exalysis > go run main.go --watch
```

Now exalysis does all automatically as soon as it detects a exercism download command in the clipboard:
- Detect and get `exercism download --uuid=...` command from clipboard
- Download exercise and load its code
- Examine code and add answer to student to clipboard

### Single exercise examination

Start it from the folder the exercise to be mentored was downloaded to. It checks code in the current directory.

```
exalysis
```

## Features

- Autodetect the exercise from the package name and the student's name from the current path.
- Run a list of tools to check the code and output the result. See supported tools below.
- Check supported exercises for special patterns and create todos, suggestions and comments based on that.
- Sometimes adds an entire block packed with knowledge about a certain topic.
- Outputs a complete answer ready to be pasted into exercism. The answer is already in the clipboard when the tool finishes.
- Provides a suggestion to the mentor whether to approve or not.

## My Typical Workflow
- Start exalysis in `watch` mode.
- Split screen 2/3 exercism, 1/3 terminal.
- Filter exercism by a certain exercise. For me mentoring the same exercise a few times is much more effective then constantly switching contexts.
- Start mentoring a solution and copy the download command.
- Wait for exalysis to run on in the terminal. It will put the answer to be pasted in the clipboard.
- Paste the answer to the student on exercism and validate what it says against the students solution.
- Adjust the answer and submit/approve.

**Warning!** The tool is **not** to be used without examining the code and trying to figure out how to 
best help the student. The output should not be used without checking it! There might be false positives 
or missing suggestions or simply too much information for the student to handle at once.

## Contribution
Any contribution is welcome as long as it complies with the code of conduct of [exercism](https://exercism.io) 
and the mentoring guidelines!

1) You can open an issue if you find a false positive that you think will happen more than once and 
should be fixed. The same with suggestions you think are missing.

2) One very simple way to contribute is to add samples for a case that is not covered yet, 
add the comment that should be shown to the student and add a test for it. That could be done for an 
already supported exercise or one that is not supported yet. The existing test structure is a good sample.

3) There are still many exercises to be implemented. If anyone feels he's up for it that would be much appreciated.
