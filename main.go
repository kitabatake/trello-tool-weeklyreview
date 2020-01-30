package main

import (
	"fmt"
	"strings"
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
	to := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, time.Local)
	from := to.AddDate(0, 0, -7)

	cards, err := fetchTrelloCards(from, to)
	if err != nil {
		fmt.Println(err)
		return
	}

	dailyCardsList := cardsToDailyCards(cards)

	out := fmt.Sprintf("%s ~ %s\n\n", from.Format("2006-01-02"), to.Format("01-02"))
	fmt.Println(out + generateMarkdown(dailyCardsList))
}

func cardsToDailyCards(cards []TrelloCard) []DailyCards {
	dailyCardsMap := make(map[string]*DailyCards)
	for _, card := range cards {

		// Why subtract one day? Because card is archived at next morning of completion on assumed trello operation.
		dayStr := card.DateLastActivity.AddDate(0, 0, -1).Format("01/02(Mon)")
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

func generateMarkdown(dailyCards []DailyCards) string {
	out := ""
	for _, dailyCard := range dailyCards {
		out += fmt.Sprintf("# %s\n\n", dailyCard.day)
		for _, card := range dailyCard.cards {
			out += fmt.Sprintf("## %s\n", card.Name)
			if len(strings.TrimSpace(card.Desc)) > 0 {
				out += "```\n" + card.Desc + "\n```\n"
			}
			out += "\n"
		}
		out += "\n---\n\n"
	}
	return out
}