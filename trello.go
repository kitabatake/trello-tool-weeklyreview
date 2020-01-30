package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"
)

type TrelloBoard struct {
	Id string `json:"id"`
}

type TrelloCard struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DateLastActivity time.Time `json:"dateLastActivity"`
	Desc string `json:"desc"`
}

type sortTrelloCards struct {
	s []TrelloCard
}

func (s sortTrelloCards)Len() int { return len(s.s) }
func (s sortTrelloCards)Less(i, j int) bool { return s.s[i].DateLastActivity.Before(s.s[j].DateLastActivity) }
func (s sortTrelloCards)Swap(i, j int) { s.s[i], s.s[j] = s.s[j], s.s[i] }

var (
	boardsUrl = "https://api.trello.com/1/members/me/boards?%s"
	cardsUrl = "https://api.trello.com/1/boards/%s/cards?%s"
)

func fetchTrelloCards(from, to time.Time) ([]TrelloCard, error) {
	boards, err := trelloApiBoards()
	if err != nil {
		return nil, err
	}

	cards := make([]TrelloCard, 0)
	for _, board := range boards {
		_cards, err := trelloApiCards(board.Id)
		if err != nil {
			return nil, err
		}
		cards = append(cards, _cards...)
	}

	filteredCards := []TrelloCard{}
	for _, card := range cards {
		l := card.DateLastActivity.Local()
		if l.After(from) && l.Before(to) {
			filteredCards = append(filteredCards, card)
		}
	}

	sort.Sort(sortTrelloCards{filteredCards})
	return filteredCards, nil
}

func trelloApiCards(bordId string) ([]TrelloCard, error) {
	urlParams := url.Values{}
	urlParams.Add("key", os.Getenv("TRELLO_API_KEY"))
	urlParams.Add("token", os.Getenv("TRELLO_TOKEN"))
	urlParams.Add("fields", "name,dateLastActivity,desc")
	urlParams.Add("filter", "closed")
	urlParams.Add("since", "2020-01-10")

	requestUrl := fmt.Sprintf(cardsUrl, bordId, urlParams.Encode())
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error occurred on api request proccess. response's http status is %s\n", resp.Status)
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

func trelloApiBoards() ([]TrelloBoard, error) {
	urlParams := url.Values{}
	urlParams.Add("key", os.Getenv("TRELLO_API_KEY"))
	urlParams.Add("token", os.Getenv("TRELLO_TOKEN"))
	urlParams.Add("fields", "id")
	urlParams.Add("filter", "open")

	requestUrl := fmt.Sprintf(boardsUrl, urlParams.Encode())
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error occurred on api request proccess. response's http status is %s\n", resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var boards []TrelloBoard
	err = json.Unmarshal(body, &boards)
	if err != nil {
		panic(err)
		return nil, err
	}

	return boards, nil
}