package epic

import (
	"time"
)

func CheckFreeGame(g Games) ([]string, []int) {
	var freeGamesTitles []string //list - games on promotion
	var num []int                //list - index of a game on promotion
	now := time.Now()

	for i := 0; i < len(g.GetAllElements()); i++ {
		if len(g.GetPromotions(i)) != 0 {
			if g.GetDiscountedPrice(i) == 0 {
				if g.GetPromotionStart(i).Before(now) && g.GetPromotionEnd(i).After(now) {
					freeGamesTitles = append(freeGamesTitles, g.GetElementTitle(i))
					num = append(num, i)
				}
			}
		}
	}
	return freeGamesTitles, num
}

func PrepMessege(g Games, titles []string, index []int) string {

	var msg string

	msg = "Get your free game for this week!\n"
	for i := 0; i < len(titles); i++ {
		if i == len(titles)-1 {
			msg += "<https://www.epicgames.com/store/pl/p/" + g.GetUrlSlug(index[i]) + "|" + titles[i] + ">"
		} else {
			msg += "<https://www.epicgames.com/store/pl/p/" + g.GetUrlSlug(index[i]) + "|" + titles[i] + ">\n"
		}

	}
	msg += "\nHave fun :)"

	return msg
}
