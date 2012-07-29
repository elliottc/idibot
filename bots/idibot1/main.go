package main

import (
	"idibot/ants"
	"io"
	"log"
)

func main() {
	var s ants.State
	err := s.Start()
	if err != nil {
		log.Panicf("Start() failed (%s)", err)
	}
	bot := NewBot(&s)
	err = s.Loop(bot, func() {
		//if you want to do other between-turn debugging things, you can do them here
	})
	if err != nil && err != io.EOF {
		log.Panicf("Loop() failed (%s)", err)
	}
}
