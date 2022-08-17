package logic

import (
	"fmt"
	"go_wordle/modules/helper"
	"go_wordle/modules/ui"
	"strings"

	"github.com/fatih/color"
)

var letters [25]int
var wordOfDay = "goods"
var userResponse = make([]LetterPair, 0)
var word string

// struct to store the status and the letter entered by the user
type LetterPair struct {
	letter string
	status int
}

func updateResponseGrid(word string) (userWon bool) {
	charArray := strings.Split(word, "")
	correctLetters := strings.Split(wordOfDay, "")
	winStatus := true
	for index, char := range charArray {
		contains := strings.Contains(wordOfDay, char)
		isAtRightIndex := correctLetters[index] == charArray[index]
		status := 0
		if isAtRightIndex {
			status = 1
		} else if contains {
			winStatus = false
			status = 2
		} else {
			winStatus = false
			status = -1
		}
		userResponse = append(userResponse, LetterPair{char, status})
	}
	return (winStatus)
}

func compareResponseWithWord(word string, letters *[25]int) {
	chars := []rune(word)
	for index, char := range chars {
		charValue := int(char)
		contains := strings.Contains(wordOfDay, string(char))
		isAtRightIndex := false
		if contains {
			isAtRightIndex = helper.CheckIfAtRightIndex(wordOfDay, string(char), index)
		}
		if isAtRightIndex {
			letters[charValue-97] = 1
		} else if contains {
			letters[charValue-97] = 2
		} else {
			letters[charValue-97] = -1
		}
	}
}

func InitializeGame() {
	var chance = 1
	userWon := false
	for chance <= 5 && !userWon {
		fmt.Printf("Enter the word\n")
		fmt.Scan(&word)
		if len(word) != 5 {
			fmt.Printf("%v is not a valid word", word)
			continue
		} else {
			userWon = updateResponseGrid(word)
			compareResponseWithWord(word, &letters)
			PrintResponseGrid()
			ui.PrintGrid(letters)
			chance += 1
		}
	}

	if userWon {
		fmt.Printf("Hurray! You Won")
	} else {
		fmt.Printf("Oops! You lost this time")
	}

}

func PrintResponseGrid() {
	count := 1
	for _, pair := range userResponse {

		if pair.status == 1 {
			fmt.Printf("%s  ", color.GreenString(pair.letter))
		} else if pair.status == 2 {
			fmt.Printf("%s  ", color.YellowString(pair.letter))
		} else if pair.status == -1 {
			fmt.Printf("%s  ", color.BlueString(pair.letter))
		}
		if count == 5 {
			fmt.Println()
			count = 0
		}
		count += 1
	}

	println()
}
