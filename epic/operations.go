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
