package main

import (
	"fmt"

	"github.com/fatih/color"
)

var letters [25]int

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
	printGrid()
}

func getLetters(word string, letters *[25]int) {
	chars := []rune(word)
	for _, char := range chars {
		charValue := int(char)
		letters[charValue-97] = 1
	}
}

func printGrid() {
	count := 1
	for index, status := range letters {
		char := string(index + 97)
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
