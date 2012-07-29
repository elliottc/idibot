package main

import (
	"idibot/ants"
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
	dirs := []ants.Direction{ants.North, ants.East, ants.South, ants.West}

	for loc, ant := range s.Map.Ants {
		if ant != ants.MY_ANT {
			continue
		}

		// try each direction in a random order
		p := rand.Perm(4)
		for _, i := range p {
			d := dirs[i]

			loc2 := s.Map.Move(loc, d)
			if s.Map.SafeDestination(loc2) {
				s.IssueOrderLoc(loc, d)
				// there's also an s.IssueOrderRowCol if you don't have a Location handy
				break
			}
		}
	}

	// returning an error will halt the whole program!
	return nil
}
