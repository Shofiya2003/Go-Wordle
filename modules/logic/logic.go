package logic

import (
	"fmt"
	"go_wordle/modules/helper"
	"go_wordle/modules/ui"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var letters [25]int
var wordOfDay = "goods"
var userResponse = make([]LetterPair, 0)
var word string
var wg sync.WaitGroup

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
	playAgain := true
	for playAgain {
		if len(stockWords) == 0 {
			getWord()
		}
		wordOfDay = stockWords[0]
		stockWords = stockWords[1:len(stockWords)]
		var chance = 1
		var count = 1
		userWon := false
		for chance <= 5 && !userWon {
			if len(stockWords) < 3 {
				wg.Add(count)
				count++
				go getWord()
			}
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

		playAgain = concludeGame(userWon)
		userResponse = make([]LetterPair, 0)
	}
	fmt.Printf("here the game ends")
}

func concludeGame(userWon bool) bool {
	if userWon {
		fmt.Printf("\nHurray! You Won\n")

	} else {
		fmt.Printf("\nOops! You lost \n Correct word is %v\n", wordOfDay)
	}
	printFinalResult()
	fmt.Printf("type 'Y' to play again\n")
	doesUserWantToPlayAgain := ""
	fmt.Scan(&doesUserWantToPlayAgain)
	if doesUserWantToPlayAgain != "Y" && doesUserWantToPlayAgain != "y" {
		return false
	}
	return true
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

func printFinalResult() {
	count := 1
	for _, pair := range userResponse {

		if pair.status == 1 {
			fmt.Printf("🟩 ")
		} else if pair.status == 2 {
			fmt.Printf("🟨 ")
		} else if pair.status == -1 {
			fmt.Printf("⬜ ")
		}
		if count == 5 {
			fmt.Printf("\n\n")
			count = 0
		}
		count += 1
	}

	println()
}
