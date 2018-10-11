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
Start it from the folder the exercise to be mentored was downloaded to. It checks code in the current directory.

```
exalysis
```

**What it does**
- Autodetect the exercise from the package name and the student's name from the current path.
- Run a list of tools to check the code and output the result. See supported tools below.
- Check supported exercises for special patterns and create todos, suggestions and comments based on that.
- Sometimes adds an entire block packed with knowledge about a certain topic.
- Outputs a complete answer ready to be pasted into exercism. The answer is already in the clipboard when the tool finishes.
- Provides a suggestion to the mentor whether to approve or not.

**My Typical Workflow**
- Filter exercism by a certain exercise. For me mentoring the same exercise a few times is much more effective then constantly switching contexts.
- Start mentoring a solution, copy the download link and execute it in the terminal.
- Copy the path from the output of `exercism download` and switch to the directory.
- Run exalysis. It will put the answer to be pasted in the clipboard.
- Paste the answer to the student on exercism and validate what it says against the students solution.
- Adjust the answer as I see fit and submit/approve.

**Warning!** The tool is **not** to be used without examining the code and trying to figure out how to 
best help the student. The output should not be used without checking it! There might be false positives 
or missing suggestions or simply too much information to handle at once.

## Contribution
Any contribution is welcome as long as it complies with the code of conduct of [exercism](https://exercism.io) 
and the mentoring guidelines! Just send a PR ;)

If you find a false positive that you think will happen more than once and should be fixed feel free to open 
an issue. The same with suggestions you think are missing.

There are still many exercises to be implemented. If anyone feels he's up for it that would be much appreciated.
