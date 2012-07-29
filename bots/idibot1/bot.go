package main

import (
	"idibot/ants"
	"math"
	"math/rand"
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
	for loc1, ant1 := range s.Map.Ants {
		// Only interested in my ants for now.
		if ant1 != ants.MY_ANT {
			continue
		}

		// Create scores for each direction the ant can move.
		scores := make([]float64, len(dirs))

		// Update the move scores for friendly ant proximity.
		for loc2, _ := range s.Map.Ants {
			if loc1 != loc2 {
				for i, dir := range dirs {
					newLoc := s.Map.Move(loc1, dir)      // Move primary ant in this direction
					dist := locDistance(s, newLoc, loc2) // Distance from friendly ant
					scores[i] += friendlyScore / (1 + dist)
				}
			}
		}

		// Update the move scores for food proximity.
		for loc2, _ := range s.Map.Food {
			for i, dir := range dirs {
				newLoc := s.Map.Move(loc1, dir)      // Move primary ant in this direction
				dist := locDistance(s, newLoc, loc2) // Distance from food
				scores[i] += foodScore / (1 + dist)
			}
		}

		// Update the move scores for hill proximity.
		for loc2, hill := range s.Map.Hills {
			for i, dir := range dirs {
				newLoc := s.Map.Move(loc1, dir)      // Move primary ant in this direction
				dist := locDistance(s, newLoc, loc2) // Distance from hill
				if hill == ants.MY_HILL {
					scores[i] += myHillScore / (1 + dist)
				} else {
					scores[i] += enemyHillScore / (1 + dist)
				}
			}
		}

		// Make the scores all positive.
		for i := 0; i < len(scores); i++ {
			scores[i] = math.Pow(10, scores[i])
		}

		// Randomly select a move using scores as weights.
		var scoreSum float64
		for i, dir := range dirs {
			if dir == ants.NoMovement || s.Map.SafeDestination(s.Map.Move(loc1, dir)) {
				scoreSum += scores[i]
			}
		}
		r := rand.Float64() * scoreSum // [0.0, scoreSum)
		var scoreSum2 float64
		var index int
		for i, dir := range dirs {
			if dir == ants.NoMovement || s.Map.SafeDestination(s.Map.Move(loc1, dir)) {
				scoreSum2 += scores[i]
				if r <= scoreSum2 {
					index = i
					break
				}
			}
		}
		move := dirs[index]
		if move != ants.NoMovement {
			s.IssueOrderLoc(loc1, move)
		}
	}

	// returning an error will halt the whole program!
	return nil
}

func locDistance(s *ants.State, loc1, loc2 ants.Location) float64 {
	row1, col1 := s.Map.FromLocation(loc1)
	row2, col2 := s.Map.FromLocation(loc2)
	return rowColDistance(row1, col1, row2, col2)
}

func rowColDistance(row1, col1, row2, col2 int) float64 {
	rDist := row1 - row2
	colDist := col1 - col2
	return math.Sqrt(float64(rDist*rDist + colDist*colDist))
}
