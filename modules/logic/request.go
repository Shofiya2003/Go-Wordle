package logic

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fatih/color"
)

var stockWords []string

// get word from http req to api
func getWord() int {
	resp, err := http.Get("https://random-word-api.herokuapp.com/word?length=5")
	//check if error is nil
	if err != nil {
		color.RedString("Error: Try again later")
		return -1
	}
	//defer is executed when the surrounding function is returned
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.RedString("Error: Try again later")
		return -1
	}
	var data [1]string
	json.Unmarshal(body, &data)
	stockWords = append(stockWords, data[0])
	if len(stockWords) != 1 {
		wg.Done()
	}
	return 1
}
