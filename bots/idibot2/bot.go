package main

import (
	"idibot/ants"
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

// TODO: duplicated
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

	// returning an error will halt the whole program!
	return nil
}
