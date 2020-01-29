package main

import (
	"fmt"
	"time"
)

type DailyCards struct {
	day string
	cards []TrelloCard
}

func (dcs DailyCards) String() string {
	return fmt.Sprintf("%s:\n%v\n---\n", dcs.day, dcs.cards)
}

func main () {
	now := time.Now()
	to := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	from := to.AddDate(0, 0, -7)

	cards, err := fetchTrelloCards(from, to)
	if err != nil {
		fmt.Println(err)
		return
	}

	dailyCardsList := cardsToDailyCards(cards)

	fmt.Println("count is ", len(cards))
	for _, dcs := range dailyCardsList {
		fmt.Printf("%s\n", dcs)
	}
}

func cardsToDailyCards(cards []TrelloCard) []DailyCards {
	dailyCardsMap := make(map[string]*DailyCards)
	for _, card := range cards {
		dayStr := card.DateLastActivity.Format("2006/01/02")
		if dailyCard, ok := dailyCardsMap[dayStr]; ok {
			dailyCard.cards = append(dailyCard.cards, card)
		} else {
			dailyCardsMap[dayStr] = &DailyCards{
				day: dayStr,
				cards: []TrelloCard{card},
			}
		}
	}

	ret := make([]DailyCards, 0)
	for _, v := range dailyCardsMap {
		ret = append(ret, *v)
	}
	return ret
}