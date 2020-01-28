package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type TrelloCard struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DateLastActivity time.Time `json:"dateLastActivity"`
	//Desc string `json:"desc"`
}

var (
	boardsUrl = "https://api.trello.com/1/boards/qoy6rgiP/cards"
)

func fetchTrelloCards(from, to time.Time) ([]TrelloCard, error) {
	urlParams := url.Values{}
	urlParams.Add("key", os.Getenv("TRELLO_API_KEY"))
	urlParams.Add("token", os.Getenv("TRELLO_TOKEN"))
	urlParams.Add("fields", "name,dateLastActivity,desc")
	urlParams.Add("filter", "closed")
	urlParams.Add("since", "2020-01-10")

	resp, err := http.Get(boardsUrl + "?" + urlParams.Encode())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var cards []TrelloCard
	err = json.Unmarshal(body, &cards)
	if err != nil {
		return nil, err
	}

	filteredCards := []TrelloCard{}
	for _, card := range cards {
		l := card.DateLastActivity.Local()
		if l.After(from) && l.Before(to) {
			filteredCards = append(filteredCards, card)
		}
	}

	return filteredCards, nil
}
