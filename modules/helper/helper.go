package helper

func CheckIfAtRightIndex(wordOfDay string, char string, requiredIndex int) (res bool) {
	chars := []rune(wordOfDay)
	for index, element := range chars {
		if string(element) == char && requiredIndex == index {
			return true
		}
	}
	return false
}
