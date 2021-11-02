package main

import (
	"log"
	"os"
	"stankryj/WorkEpicProject/epic"
	"time"
)

func main() {
	epicURL := "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL"
	slackURL, exists := os.LookupEnv("SLACKURL")
	if exists == false {
		log.Fatal("Didn't find environment variable named SLACKURL")
	}
	var oldGame string

	for {
		epic.EpicWeekly(epicURL, slackURL, &oldGame)
		time.Sleep(time.Minute * 60)
	}
}
