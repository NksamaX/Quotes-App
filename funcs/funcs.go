package funcs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quote/types"

	// "github.com/anaskhan96/soup"
	"github.com/fatih/color"
)

func Get_Quote() {
	c := color.New(color.Bold).Add(color.BgBlue).PrintfFunc()
	res, err := http.Get("https://animechan.vercel.app/api/random")
	if err != nil {
		c("Error: %v", err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	var qt types.Quote
	json.Unmarshal(body, &qt)
	c("%v\n   - %v", qt.Quote, qt.Character)
}

func TestQuote() {
	res, err := http.Get("https://randomwordgenerator.com/json/inspirational-quote_ws.json")
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	var qt types.Qt
	json.Unmarshal(body, &qt)
	fmt.Println(qt)

}
