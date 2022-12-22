package hangman

import (
	"strings"

	"golang.org/x/exp/slices"
)

type Game struct {
	State        string
	Letters      []string
	FoundLetters []string
	UsedLetters  []string
	TurnsLeft    int
}

func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(found); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		TurnsLeft:    turns,
	}
	return g
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)

	if slices.Contains(g.UsedLetters, guess) {
		g.State = "alreadyGuessed"
	} else if slices.Contains(g.Letters, guess) {
		g.State = "goodGuess"
		g.RevealLetter(guess)

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.LoseTurn(guess)

		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, v := range g.Letters {
		if v == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	return slices.Equal(letters, foundLetters)
}

func (g *Game) LoseTurn(guess string) {
	g.TurnsLeft -= 1
	g.UsedLetters = append(g.UsedLetters, guess)
}
