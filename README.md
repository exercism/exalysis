# Exalysis

[![Build Status](https://travis-ci.org/tehsphinx/exalysis.svg?branch=master)](https://travis-ci.org/tehsphinx/exalysis)

Exalysis is a tool designed to help mentors of the [Exercism](https://Exercism.io) Go track. It will automatically run the tests on a student's solution, check `gofmt` and `golint`, and make some helpful suggestions for the student based on static analysis of the code for common errors and patterns (see [sample Exalysis output](sample.md)).

It's not perfect: you should check Exalysis's suggestions before sending them to the student, and add your own notes if necessary, but it can dramatically speed up the process of mentoring by automating these repetitive tasks.

## Status

This is a proof of concept. I use it to mentor, and enhance/extend it as I go.

***Exercise implementations:***
- Two Fer
- Hamming
- Raindrops
- Scrabble Score
- Isogram
- Difference of Squares
- Luhn

Exalysis will do its format, lint, and test checks on any solution, but specific suggestions have so far been implemented for only the exercises above. If you'd like to add support for a new exercise, please submit a PR! (See 'Contributions' below.)

## Installation

```
go get github.com/tehsphinx/exalysis
go install github.com/tehsphinx/exalysis/cmd/exalysis
```

## Usage

### Watch mode

To start Exalysis in watch mode, run this command:

```
exalysis -watch
```

When you copy an Exercism download command to the clipboard, Exalysis will detect this and do the following:

- Download exercise and run tests, lint checks, and so on
- Analyse the code and make suggestions for improvement

The output of Exalysis is automatically copied to the clipboard, so all you need to do is paste into the comment box on the Exercism site and review your answer.

### Single exercise examination

With no arguments, the `exalysis` command will analyse the solution in the current directory and write its suggestions to the standard output.

```
exalysis
```

## Features

- Auto-detect the exercise from the package name and the student's name from the current path.
- Run a list of tools to check the code and output the result. See supported tools below.
- Check supported exercises for special patterns and create todos, suggestions and comments based on that.
- Sometimes adds an entire block packed with knowledge about a certain topic.
- Outputs a complete answer ready to be pasted into Exercism. The answer is already in the clipboard when the tool finishes.
- Provides a suggestion to the mentor whether to approve or not.

## My Typical Workflow

1. Start exalysis in `-watch` mode.
1. Split screen 2/3 Exercism, 1/3 terminal.
1. Filter Exercism by a certain exercise. For me mentoring the same exercise a few times is much more effective then constantly switching contexts.
1. Start mentoring a solution and copy the download command.
1. Wait for Exalysis to run on in the terminal. It will put the answer to be pasted in the clipboard.
1. Paste the answer to the student on Exercism and validate what it says against the student's solution.
1. Adjust the answer and submit/approve.

**Warning!** The tool is **not** to be used without examining the code and trying to figure out how to best help the student. The output should not be used without checking it! There might be false positives or missing suggestions or simply too much information for the student to handle at once.

## Contribution
Any contribution is welcome as long as it complies with the code of conduct of [Exercism](https://Exercism.io) and the mentoring guidelines.

- You can open an issue if you find a false positive that you think will happen more than once and should be fixed. The same with suggestions you think are missing.

- One very simple way to contribute is to add samples for a case that is not covered yet, add the comment that should be shown to the student and add a test for it. That could be done for an already supported exercise or one that is not supported yet. The existing test structure is a good sample.

- There are still many exercises to be implemented. If you're up for implementing a new exercise, that would be much appreciated.

## Discussion

There's a channel on the Exercism Team Slack for discussion, development, and questions about Exalysis: [#track-go-exalysis](https://exercism-team.slack.com/messages/CE6EMAFEZ).

If you're not already on this Slack, email mentoring@exercism.io to request an invite.