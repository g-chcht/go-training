package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"t", "e", "s", "t"}
	guess := "t"
	hasLetter := letterInWord(word, guess)

	if !hasLetter {
		t.Errorf("Word: %s, Guess: %s, hasLetter: %v", word, guess, hasLetter)
	}
}

func TestNotLetterInWord(t *testing.T) {
	word := []string{"t", "e", "s", "t"}
	guess := "z"
	hasLetter := letterInWord(word, guess)

	if hasLetter {
		t.Errorf("Word: %s, Guess: %s, hasLetter: %v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Error("word is empty !")
	}
}

