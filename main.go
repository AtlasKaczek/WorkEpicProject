package main

import (
	"fmt"
	"stankryj/WorkEpicProject/epic"
)

func main() {
	epicUrl := "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL"

	freeGame, err := epic.ParseJson(epicUrl)
	if err != nil {
		fmt.Printf("Main 1: An error occured: %v", err)
	}

	var titles []string
	var index []int
	titles, index = epic.CheckFreeGame(freeGame)
	fmt.Printf("Title: \"%v\" Index: %d\n", titles[0], index[0])

	fmt.Println(epic.PrepMessege(freeGame, titles, index))
}
