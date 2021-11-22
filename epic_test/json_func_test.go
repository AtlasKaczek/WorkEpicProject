package epictest

import (
	"stankryj/WorkEpicProject/epic"
	"testing"
)

func TestGetJSON(t *testing.T) {
	body, err := epic.GetJSON("https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL")
	if err != nil {
		t.Fatal(err)
	}
	if body == nil {
		t.Errorf("body is nil")
	}
	if len(*body) == 0 {
		t.Errorf("body is empty")
	}
}

func TestParseJSON(t *testing.T) {
	freeGame, err := epic.ParseJSON("https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL")
	if err != nil {
		t.Fatal(err)
	}
	if freeGame.GetProductSlug(0) == "" {
		t.Errorf("Games object is empty")
	}
}
