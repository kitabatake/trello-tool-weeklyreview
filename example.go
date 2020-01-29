package main

import (
	"fmt"
)

func main () {
	exampleBoards()
}

func exampleBoards() {
	boards, err := trelloApiBoards()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(boards)
}