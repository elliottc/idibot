package main

import (
	"idibot/ants"
	"idibot/lib"
	"math"
)

type Bot struct {
}

// NewBot creates a new instance of your bot
func NewBot(s *ants.State) ants.Bot {
	bot := &Bot{
	// do any necessary initialization here
	}

	return bot
}

func myAnts(allAnts map[ants.Location]ants.Item) []ants.Location {
	a := make([]ants.Location, 0)
	for loc, item := range allAnts {
		if item == ants.MY_ANT {
			a = append(a, loc)
		}
	}

	return a
}

// DoTurn is where you should do your bot's actual work.
func (bot *Bot) DoTurn(s *ants.State) error {
	const (
		friendlyScore  = -1
		foodScore      = 3
		myHillScore    = -2
		enemyHillScore = 1
	)

	// Score the possible moves for each ant.
	dirs := []ants.Direction{ants.North, ants.East, ants.South, ants.West, ants.NoMovement}
	for _, loc1 := range myAnts(s.Map.Ants) {
		// Create moves and scores for each direction the ant can move.
		moves := make([]ants.Direction, 0)
		scores := make([]float64, 0)
		for i, dir := range dirs {
			// Don't include this move if it is not legal.
			newLoc := s.Map.Move(loc1, dir)
			if !s.Map.SafeDestination(newLoc) {
				continue
			}

			var score float64
			row1, col1 := s.Map.FromLocation(newLoc)

			// Update the move scores for friendly ant proximity.
			for _, loc2 := range myAnts(s.Map.Ants) {
				if loc1 != loc2 {
					row2, col2 := s.Map.FromLocation(loc2)
					dist := lib.ManhattanDistance(row1, col1, row2, col2, s.Map.Rows, s.Map.Cols)
					score += friendlyScore / float64(1+dist)
				}
			}

			// Update the move scores for food proximity.
			for loc2, _ := range s.Map.Food {
				row2, col2 := s.Map.FromLocation(loc2)
				dist := lib.ManhattanDistance(row1, col1, row2, col2, s.Map.Rows, s.Map.Cols)
				score += foodScore / float64(1+dist)
			}

			// Update the move scores for hill proximity.
			for loc2, hill := range s.Map.Hills {
				row2, col2 := s.Map.FromLocation(loc2)
				dist := lib.ManhattanDistance(row1, col1, row2, col2, s.Map.Rows, s.Map.Cols)
				if hill == ants.MY_HILL {
					score += myHillScore / float64(1+dist)
				} else {
					score += enemyHillScore / float64(1+dist)
				}
			}

			moves = append(moves, dirs[i])
			scores = append(scores, score)
		}

		// Make the scores all positive.
		for i := 0; i < len(scores); i++ {
			scores[i] = math.Exp(scores[i])
		}

		// Randomly select a move using scores as weights.
		if len(scores) > 0 {
			index := lib.RandSelect(scores)
			move := moves[index]
			if move != ants.NoMovement {
				s.IssueOrderLoc(loc1, move)
			}
		}
	}

	// returning an error will halt the whole program!
	return nil
}
