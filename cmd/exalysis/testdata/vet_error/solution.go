// Package tournament keeps tab on USPL
package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	firstTeamName = iota
	secondTeamName
	matchResult
)

const poolSize = 4

var poolA = [poolSize]string{
	"Allegoric Alaskians",
	"Blithering Badgers",
	"Courageous Californians",
	"Devastating Donkeys",
}

type score struct {
	teamName      string
	matchesPlayed int
	won           int
	drawn         int
	lost          int
	points        int
}

// Tally reads raw data from input and outpts tallied results
func Tally(input io.Reader, output io.Writer) error {

	var scoreCard [poolSize]*score
	var nameToIndex = map[string]int{}

	// Create score card to keep track of result
	for i, name := range poolA {
		card := new(score)
		card.teamName = name
		scoreCard[i] = card
		nameToIndex[name] = i
	}

	// INPUT

	var resultsTable strings.Builder

	p := make([]byte, 100)
	for {
		n, err := input.Read(p)
		resultsTable.Write(p[:n])
		if io.EOF == err {
			break
		}
	}

	results := strings.Split(resultsTable.String(), "\n")

	for _, resultLine := range results {
		details := strings.Split(resultLine, ";")

		// VALIDATE

		switch {
		// Ignore commented lines
		case strings.HasPrefix(resultLine, "#"):
			continue
			// Ignore emplty lines
		case len(details) == 1 && len(details[0]) == 0:
			continue
			// Not fully formed
		case len(details) < 3:
			return errors.New("Invalid input")
			continue
		}

		// PROCESS

		var firstTeam, secondTeam *score
		var index int
		var teamValid bool

		// Get first team stats
		teamName := details[firstTeamName]
		index, teamValid = nameToIndex[teamName]
		if !teamValid {
			return errors.New("Invalid team found")
		}
		firstTeam = scoreCard[index]

		// Get second team stats
		teamName = details[secondTeamName]
		index, teamValid = nameToIndex[teamName]
		if !teamValid {
			return errors.New("Invalid team found")
		}
		secondTeam = scoreCard[index]

		// Update scores.
		firstTeam.matchesPlayed++
		secondTeam.matchesPlayed++

		switch details[matchResult] {
		case "win":
			firstTeam.won++
			secondTeam.lost++
			firstTeam.points += 3

		case "loss":
			firstTeam.lost++
			secondTeam.won++
			secondTeam.points += 3
		case "draw":
			firstTeam.drawn++
			secondTeam.drawn++
			firstTeam.points++
			secondTeam.points++
		default:
			return errors.New("Invalid result found")
		}
	}

	// SORT RESULTS

	sortedCard := scoreCard[:]
	sort.Slice(sortedCard, func(i, j int) bool {
		if sortedCard[i].points == sortedCard[j].points {
			return sortedCard[i].teamName < sortedCard[j].teamName
		}
		return sortedCard[i].points > sortedCard[j].points
	})

	// OUTPUT

	fmt.Fprintf(output, "%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")

	for _, v := range sortedCard {
		fmt.Fprintf(output, "%-31s|%3d |%3d |%3d |%3d |%3d\n",
			v.teamName, v.matchesPlayed, v.won, v.drawn, v.lost, v.points)
	}

	return nil
}
