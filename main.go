package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
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
	printGrid(letters)
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

func printResponseGrid() {
	count := 1
	for _, element := range userLetters {
		char := string(element)
		status := letters[element-97]
		if status == 1 {
			fmt.Printf("%s  ", color.YellowString(char))
		} else {
			fmt.Printf("%s  ", char)
		}
		if count == 5 {
			fmt.Println()
			count = 0
		}
		count += 1
	}
}

func printGrid(grid [25]int) {
	count := 1
	for index, status := range grid {
		char := string(index + 97)
		if status == 1 {
			fmt.Printf("%s  ", color.GreenString(char))
		} else if status == 2 {
			fmt.Printf("%s  ", color.YellowString(char))
		} else if status == -1 {
			fmt.Printf("%s  ", color.BlueString(char))
		} else {
			fmt.Printf("%s  ", char)
		}
		if count == 5 {
			fmt.Println()
			count = 0
		}
		count += 1
	}
}
