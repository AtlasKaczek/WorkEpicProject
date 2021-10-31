package main

import (
	"fmt"
	"stankryj/WorkEpicProject/epic"
)

func main() {
	epicUrl := "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL"
	slackUrl := "https://hooks.slack.com/services/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"

	freeGame, err := epic.ParseJson(epicUrl)
	if err != nil {
		fmt.Printf("Main 1: An error occured: %v", err)
	}

	var titles []string
	var index []int
	titles, index = epic.CheckFreeGame(freeGame)

	_, err = epic.SendSlackMessege(epic.PrepMessege(freeGame, titles, index), slackUrl)
	if err != nil {
		fmt.Printf("Main 2: An error occured: %v", err)
	}
}
