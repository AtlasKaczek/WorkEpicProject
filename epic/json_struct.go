package epic

import (
	"time"
)

//Structure
type Games struct {
	Data Data `json:"data"`
}

type Data struct {
	Catalog Catalog `json:"Catalog"`
}

type Catalog struct {
	SearchStore SearchStore `json:"searchStore"`
}

type SearchStore struct {
	Elements []Element `json:"elements"`
}

type Element struct {
	Title       string    `json:"title"`
	Namespace   string    `json:"namespace"`
	Description string    `json:"description"`
	ProductSlug string    `json:"productSlug"`
	URLSlug     string    `json:"urlSlug"`
	Price       Price     `json:"price"`
	Promotions  Promotion `json:"promotions"`
}

type FMTPrice struct {
	OriginalPrice     string `json:"originalPrice"`
	DiscountPrice     string `json:"discountPrice"`
	IntermediatePrice string `json:"intermediatePrice"`
}

type TotalPrice struct {
	DiscountPrice   int          `json:"discountPrice"`
	OriginalPrice   int          `json:"originalPrice"`
	VoucherDiscount int          `json:"voucherDiscount"`
	Discount        int          `json:"discount"`
	CurrencyCode    string       `json:"currencyCode"`
	CurrencyInfo    CurrencyInfo `json:"currencyInfo"`
	FMTPrice        FMTPrice     `json:"fmtPrice"`
}

type Price struct {
	TotalPrice TotalPrice `json:"totalPrice"`
}

type CurrencyInfo struct {
	Decimals int `json:"decimals"`
}

type Promotion struct {
	PromotionalOffers         []PromotionalOffers `json:"promotionalOffers"`
	UpcomingPromotionalOffers []PromotionalOffer  `json:"upcomingPromotionalOffers"`
}

type PromotionalOffers struct {
	PromotionalOffer []PromotionalOffer `json:"promotionalOffers"`
}

type PromotionalOffer struct {
	StartDate       time.Time       `json:"startDate"`
	EndDate         time.Time       `json:"endDate"`
	DiscountSetting DiscountSetting `json:"discountSetting"`
}

type DiscountSetting struct {
	DiscountType       string `json:"discountType"`
	DiscountPercentage int    `json:"discountPercentage"`
}

//Useful getters
func (g Games) GetAllElements() []Element {
	return g.Data.Catalog.SearchStore.Elements
}

func (g Games) GetElementTitle(i int) string {
	return g.Data.Catalog.SearchStore.Elements[i].Title
}

func (g Games) GetPromotions(i int) []PromotionalOffers {
	return g.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers
}

func (g Games) GetPromotionStart(i int) time.Time {
	return g.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers[0].PromotionalOffer[0].StartDate
}

func (g Games) GetPromotionEnd(i int) time.Time {
	return g.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers[0].PromotionalOffer[0].EndDate
}

func (g Games) GetDiscountedPrice(i int) int {
	return g.Data.Catalog.SearchStore.Elements[i].Price.TotalPrice.DiscountPrice
}

func (g Games) GetUrlSlug(i int) string {
	return g.Data.Catalog.SearchStore.Elements[i].URLSlug
}
