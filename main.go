package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	//"time"
)

type TrelloCard struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DateLastActivity time.Time `json:"dateLastActivity"`
}

func main () {
	values := url.Values{}
	values.Add("key", os.Getenv("TRELLO_API_KEY"))
	values.Add("token", os.Getenv("TRELLO_TOKEN"))
	values.Add("fields", "name,dateLastActivity")
	values.Add("filter", "closed")
	values.Add("since", "2020-01-28")

	endpoint := "https://api.trello.com/1/boards/qoy6rgiP/cards"
	resp, err := http.Get(endpoint + "?" + values.Encode())
	if err != nil {
		fmt.Println("error!")
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("status :" + resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var cards []TrelloCard
	err = json.Unmarshal(body, &cards)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cards)
}
