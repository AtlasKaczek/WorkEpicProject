package main

import (
	"stankryj/WorkEpicProject/epic"
	"time"
)

func main() {
	epicUrl := "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL"
	slackUrl := "https://hooks.slack.com/services/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"

	var oldGame string

	for {
		epic.EpicWeekly(epicUrl, slackUrl, &oldGame)
		time.Sleep(time.Minute * 60)
	}
}
