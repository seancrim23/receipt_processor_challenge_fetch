package models

import (
	"strconv"
)

/*
	{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
*/
//build out more
type Receipt struct {
	Id           string `json:"id"`
	Retailer     string `json:"retailer" validate:"min=1,regexp=^[\\w\\s\\-&]+$"`
	PurchaseDate string `json:"purchaseDate" validate:"min=1,regexp=^\\d{4}-\\d{2}-\\d{2}$"`
	PurchaseTime string `json:"purchaseTime" validate:"min=1,regexp=^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$"`
	Items        []Item `json:"items" validate:"min=1"`
	Total        string `json:"total" validate:"min=1,regexp=^\\d+\\.\\d{2}$"`
}

func (r *Receipt) CalculateTotal() string {
	total := 0.00
	for _, v := range r.Items {
		priceFloat, _ := strconv.ParseFloat(v.Price, 64)
		total += priceFloat
	}
	return strconv.FormatFloat(total, 'f', -1, 64)
}

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"min=1,regexp=^[\\w\\s\\-]+$"`
	Price            string `json:"price" validate:"min=1,regexp=^\\d+\\.\\d{2}$"`
}
