package main

import (
	"encoding/json"
	"fmt"
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
	boardsUrl = "https://api.trello.com/1/boards/%s/cards?%s"
)

func fetchTrelloCards(from, to time.Time) ([]TrelloCard, error) {
	cards, err := trelloApiCards("qoy6rgiP")
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

func trelloApiCards(bordId string) ([]TrelloCard, error) {
	urlParams := url.Values{}
	urlParams.Add("key", os.Getenv("TRELLO_API_KEY"))
	urlParams.Add("token", os.Getenv("TRELLO_TOKEN"))
	urlParams.Add("fields", "name,dateLastActivity,desc")
	urlParams.Add("filter", "closed")
	urlParams.Add("since", "2020-01-10")

	requestUrl := fmt.Sprintf(boardsUrl, bordId, urlParams.Encode())
	resp, err := http.Get(requestUrl)
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

	return cards, nil
}