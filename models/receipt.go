package models

type Receipt struct {
	Id           string `json:"id"`
	Retailer     string `json:"retailer" validate:"min=1,regexp=^[\\w\\s\\-&]+$"`
	PurchaseDate string `json:"purchaseDate" validate:"min=1,regexp=^\\d{4}-\\d{2}-\\d{2}$"`
	PurchaseTime string `json:"purchaseTime" validate:"min=1,regexp=^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$"`
	Items        []Item `json:"items" validate:"min=1"`
	Total        string `json:"total" validate:"min=1,regexp=^\\d+\\.\\d{2}$"`
}

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"min=1,regexp=^[\\w\\s\\-]+$"`
	Price            string `json:"price" validate:"min=1,regexp=^\\d+\\.\\d{2}$"`
}
