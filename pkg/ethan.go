package ethan

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	MaxNumPlayers    = 50
	MaxStartingChips = 1000
)

type EthanGame struct {
	// json lib by default doesn't print out things
	// that are lowercase unexposed so i capitalized them all here
	// even though outside of PrintFinalSummary it's not necessary
	Ethan          int
	Players        []int
	EthanEyes      bool
	AutoPlay       bool
	CurrentPlayer  int
	outputFunction func(string)
	TurnCount      int
	StartingChips  int
}

func InitializeEthan(startingChips int, numPlayers int, ethanEyes bool, autoPlay bool, whereToOutput string) (EthanGame, error) {
	game := EthanGame{
		EthanEyes:     ethanEyes,
		AutoPlay:      autoPlay,
		Ethan:         0,
		CurrentPlayer: 0,
		TurnCount:     1,
	}
	// basic validation
	if numPlayers < 1 {
		return game, fmt.Errorf("number of Players cannot be less than 1\n")
	}
	if numPlayers > MaxNumPlayers {
		return game, fmt.Errorf("number of Players cannot exceed %d\n", MaxNumPlayers)
	}
	if startingChips < 1 {
		return game, fmt.Errorf("cannot have fewer than 1 starting chip\n")
	}
	if startingChips > MaxStartingChips {
		return game, fmt.Errorf("cannot have more than %d starting chips\n", MaxStartingChips)
	}
	game.StartingChips = startingChips
	game.Players = make([]int, numPlayers)
	for idx, _ := range game.Players {
		game.Players[idx] = startingChips
	}
	// TODO: multiple output functions would be nice but for now, let's just log it all out
	game.outputFunction = func(outputStr string) {
		log.Println(outputStr)
	}
	return game, nil
}

func (game *EthanGame) StartGame() error {
	rand.Seed(time.Now().Unix())
	for {
		// there always has to be at least one turn
		skipped := game.executeTurn()
		if !skipped {
			if game.checkWinCondition() || game.checkLoseCondition() {
				break
			}
		}
		// if we're not done, increment the player
		game.incrementPlayerAndTurn(skipped)
	}
	return nil
}

func (game *EthanGame) incrementPlayerAndTurn(skipped bool) {
	game.CurrentPlayer = (game.CurrentPlayer + 1) % len(game.Players)
	if !skipped {
		game.TurnCount += 1
	}
}

func (game *EthanGame) checkWinCondition() bool {
	// if the current player has all of teh chips, done deal.
	if game.Players[game.CurrentPlayer] == (game.StartingChips * len(game.Players)) {
		game.outputFunction(fmt.Sprintf("player %d has all of the chips! player %d wins!",
			game.CurrentPlayer, game.CurrentPlayer))
		return true
	}
	return false
}

func (game *EthanGame) checkLoseCondition() bool {
	for _, val := range game.Players {
		if val != 0 {
			// somebody has chips, game isn't over
			return false
		}
	}
	game.outputFunction("no Players have any chips left, Ethan wins, everybody loses, gg")
	return true
}

func (game *EthanGame) executeTurn() bool {
	if game.Players[game.CurrentPlayer] == 0 {
		// skip this player, out of chips.
		return true
	}
	game.outputFunction(fmt.Sprintf("Turn %d, player %d", game.TurnCount, game.CurrentPlayer))
	if !game.AutoPlay {
		game.outputFunction("hit enter to execute turn")
		// wait for some keyboard input
		fmt.Scanln()
	}
	die1 := (rand.Int() % 6) + 1
	die2 := (rand.Int() % 6) + 1
	game.outputFunction(fmt.Sprintf("die 1: %d -- die 2: %d", die1, die2))
	if game.EthanEyes && (die1+die2) == 2 {
		game.outputFunction(fmt.Sprintf("Ethan eyes, player %d loses it all", game.CurrentPlayer))
		game.Ethan += game.Players[game.CurrentPlayer]
		game.Players[game.CurrentPlayer] = 0
	} else if (die1 + die2) == 4 {
		// oh good a positive result
		game.outputFunction(fmt.Sprintf("rolled a 4, player %d gets all of Ethan's current chips, %d",
			game.CurrentPlayer, game.Ethan))
		game.Players[game.CurrentPlayer] += game.Ethan
		game.Ethan = 0
	} else {
		//rolled a dud
		game.outputFunction(fmt.Sprintf("rolled %d, player %d gives a chip to Ethan who now has %d",
			die1+die2, game.CurrentPlayer, game.Ethan))
		game.Players[game.CurrentPlayer] -= 1
		game.Ethan += 1
	}
	game.printPlayerStatus()
	return false
}

func (game *EthanGame) printPlayerStatus() {
	game.outputFunction(fmt.Sprintf("Ethan has %d", game.Ethan))
	for idx,val := range game.Players {
		game.outputFunction(fmt.Sprintf("player %d has %d", idx, val))
	}
}

func (game *EthanGame) PrintFinalSummary() error {
	// basically just json pretty print the final status
	outputJson, err := json.MarshalIndent(*game, "", "   ")
	if err != nil {
		return err
	}
	game.outputFunction(fmt.Sprintf("final state \n%s", outputJson))
	return nil
}
