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
}
