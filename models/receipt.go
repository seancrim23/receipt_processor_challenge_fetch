package models

import "strconv"

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
	Id           string
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        string
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
	ShortDescription string
	Price            string
}
