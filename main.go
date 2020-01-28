package main

import (
	"fmt"
	"os"
)

func main () {
	fmt.Println("yeah!" + os.Getenv("TRELLO_API_KEY"))
}
