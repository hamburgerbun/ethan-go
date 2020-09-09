package main

import (
	"flag"
	"fmt"
	ethan "github.com/ethan-go/pkg"
	"log"
)

func main() {
	ethanEyes := flag.Bool("e", false, "ethan eyes ruleset on")
	autoPlayer := flag.Bool("a", false, "autoplay on, no need to hit enter")
	startingChips := flag.Int("c", 5, fmt.Sprintf("number of chips to start with, greater than 0 but less than %d", ethan.MaxStartingChips))
	numPlayer := flag.Int("p", 5, fmt.Sprintf("number of players other than ethan, greater than 0 but less than %d", ethan.MaxNumPlayers))

	flag.Parse()

	theGame, err := ethan.InitializeEthan(*startingChips, *numPlayer, *ethanEyes, *autoPlayer, "log")
	if err != nil {
		log.Fatalf("failed to init ethan : %v\n", err)
	}

	if *autoPlayer {
		// let's start it in a go function because we can lollll not like there's a point though i guess
		// we could run multiple games at once.
		errorCh := make(chan error)
		go func() {
			errorCh <- theGame.StartGame()
		}()
		err = <- errorCh
	} else {
		err = theGame.StartGame()
	}
	if err != nil {
		log.Fatalf("failed to play ethan : %v\n", err)
	}
	err = theGame.PrintFinalSummary()
	if err != nil {
		log.Fatalf("failed to print final ethan structure : %v\n", err)
	}
}
