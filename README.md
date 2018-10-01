# Exalysis

Exalysis is a tool ment for mentors of the [exercism](https://exercism.io) go track.

## Status
This is a prove of concept. I use it to mentor and enhance/extend it as I go.

***Exercise implementations:***
- TwoFer (in progress)
- Scrabble Score 

## Installation

```
GO111MODULE=on go install github.com/tehsphinx/exalysis/cmd/exalysis
```

## Usage
Start it from the folder the exercise to be mentored was downloaded to. It checks code in the current directory.

```
exalysis
```

It will autodetect the exercise from the package name and the student's name from the current path.

**Warning!** The tool is **not** to be used without examining the code and trying to figure out how to 
best help the student. The output should not be used without checking it!   

## Contribution
Any contribution is welcome as long as it complies with the code of conduct of [exercism](https://exercism.io) 
and the mentoring guidelines! Just send a PR.