package logic

import (
	"encoding/json"
	"io"
	"net/http"
)

var stockWords []string

// get word from http req to api
func getWord() {
	resp, err := http.Get("https://random-word-api.herokuapp.com/word?length=5")
	//check if error is nil
	if err != nil {
		//handle the error here
	}
	//defer is executed when the surrounding function is returned
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var data [1]string
	json.Unmarshal(body, &data)
	stockWords = append(stockWords, data[0])
	if len(stockWords) != 1 {
		wg.Done()
	}
}
