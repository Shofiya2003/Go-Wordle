package ui

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintGrid(grid [25]int) {
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
