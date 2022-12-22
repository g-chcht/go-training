package main

import (
	"fmt"
	"os"

	"training.go/hangman/hangman"
)

func main() {

	g := hangman.New(8, "test")

	hangman.DrawWelcome()

	guess := ""

	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("An error happened: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}
