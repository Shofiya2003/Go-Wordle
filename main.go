package main

import (
	"fmt"
	"go_wordle/modules/ui"
	"strings"
)

var letters [25]int
var wordOfDay = "goods"
var userLetters []int

func main() {
	var word string
	for true {
		fmt.Printf("Enter the word\n")
		fmt.Scan(&word)
		if len(word) != 5 {
			fmt.Printf("%v is not a valid word", word)
			continue
		} else {
			break
		}
	}
	getLetters(word, &letters)
	ui.PrintGrid(letters)
}

func getLetters(word string, letters *[25]int) {
	chars := []rune(word)
	compareResponseWithWord(word, letters)
	for _, char := range chars {
		charValue := int(char)
		userLetters = append(userLetters, charValue)
	}
}

func compareResponseWithWord(word string, letters *[25]int) {
	chars := []rune(word)
	for index, char := range chars {
		charValue := int(char)
		contains := strings.Contains(wordOfDay, string(char))
		isAtRightIndex := false
		if contains {
			isAtRightIndex = checkIfAtRightIndex(string(char), index)
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

func checkIfAtRightIndex(char string, requiredIndex int) (res bool) {
	chars := []rune(wordOfDay)
	for index, element := range chars {
		if string(element) == char && requiredIndex == index {
			return true
		}
	}
	return false
}
