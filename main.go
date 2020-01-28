package main

import (
	"fmt"
	"time"
)

func main () {
	now := time.Now()
	to := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	from := to.AddDate(0, 0, -7)

	cards, err := fetchTrelloCards(from, to)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("count is ", len(cards))
	fmt.Println(cards)
}
