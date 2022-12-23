package main

import (
	"fmt"
	"os"

	"training.go/hangman/dictionary"
	"training.go/hangman/hangman"
)

func main() {

	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Something happened: %v", err)
		os.Exit(1)
	}

	g, err := hangman.New(8, dictionary.PickWorld())
	if err != nil {
		fmt.Printf("Something happened: %v", err)
		os.Exit(1)
	}

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
