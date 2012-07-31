package main

import (
	"idibot/ants"
	"io"
	"log"
)

//main initializes the state and starts the processing loop
func main() {
	var s ants.State
	err := s.Start()
	if err != nil {
		log.Panicf("Start() failed (%s)", err)
	}
	mb := NewBot(&s)
	err = s.Loop(mb, func() {
		//if you want to do other between-turn debugging things, you can do them here
	})
	if err != nil && err != io.EOF {
		log.Panicf("Loop() failed (%s)", err)
	}
}
