package main

import (
	"stankryj/WorkEpicProject/epic"
	"time"
)

func main() {
	epicUrl := "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL"
	slackUrl := "https://hooks.slack.com/services/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"
	for {
		epic.EpicWeekly(epicUrl, slackUrl)
		time.Sleep(time.Minute * 60)
	}
}
