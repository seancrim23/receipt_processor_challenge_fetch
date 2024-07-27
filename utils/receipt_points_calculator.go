package utils

import (
	"fmt"
	"math"
	"receipt_processor/models"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type ReceiptPointsCalculator struct {
	Rules []Rule
}

type Rule interface {
	CalculatePoints(r models.Receipt) int
}

type AlphanumericRule struct{}

// One point for every alphanumeric character in the retailer name.
func (a AlphanumericRule) CalculatePoints(r models.Receipt) int {
	points := 0
	for _, v := range r.Retailer {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			points += 1
		}
	}
	return points
}

type RoundDollarTotalRule struct{}

// 50 points if the total is a round dollar amount with no cents.
func (rd RoundDollarTotalRule) CalculatePoints(r models.Receipt) int {
	totalFloat, _ := strconv.ParseFloat(r.Total, 64)
	if totalFloat == math.Trunc(totalFloat) {
		return 50
	}
	return 0
}

type MultipleOfQuarterTotalRule struct{}

// 25 points if the total is a multiple of 0.25.
func (m MultipleOfQuarterTotalRule) CalculatePoints(r models.Receipt) int {
	totalFloat, _ := strconv.ParseFloat(r.Total, 64)
	if math.Mod(totalFloat, 0.25) == 0 {
		return 25
	}
	return 0
}

type TwoItemsRule struct{}

// 5 points for every two items on the receipt.
func (t TwoItemsRule) CalculatePoints(r models.Receipt) int {
	points := 0
	for i, _ := range r.Items {
		if i%2 != 0 {
			points += 5
		}
	}
	return points
}

type TrimmedLengthRule struct{}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
// originally using math.round but everything needs rouding up so ceil
func (t TrimmedLengthRule) CalculatePoints(r models.Receipt) int {
	points := 0
	for _, v := range r.Items {
		trimmedDesc := strings.Trim(v.ShortDescription, " ")
		if len(trimmedDesc)%3 == 0 {
			floatPrice, _ := strconv.ParseFloat(v.Price, 64)
			points += int(math.Ceil(floatPrice * 0.2))
		}
	}
	return points
}

type OddPurchaseDateRule struct{}

// 6 points if the day in the purchase date is odd.
func (t OddPurchaseDateRule) CalculatePoints(r models.Receipt) int {
	purchaseDate, err := time.Parse("2006-01-02", r.PurchaseDate)
	if err != nil {
		return 0
	}
	if purchaseDate.Day()%2 != 0 {
		return 6
	}
	return 0
}

type TimeOfPurchaseBetween2and4Rule struct{}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func (t TimeOfPurchaseBetween2and4Rule) CalculatePoints(r models.Receipt) int {
	purchaseTime, err := time.Parse(time.DateTime, r.PurchaseDate+" "+r.PurchaseTime+":00")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	twoPmDate := time.Date(purchaseTime.Year(), purchaseTime.Month(), purchaseTime.Day(), 14, 0, 0, 0, time.UTC)
	fourPmDate := time.Date(purchaseTime.Year(), purchaseTime.Month(), purchaseTime.Day(), 16, 0, 0, 0, time.UTC)
	if purchaseTime.After(twoPmDate) && purchaseTime.Before(fourPmDate) {
		return 10
	}
	return 0
}

func NewReceiptPointsCalculator() *ReceiptPointsCalculator {
	rules := []Rule{
		AlphanumericRule{},
		RoundDollarTotalRule{},
		MultipleOfQuarterTotalRule{},
		TwoItemsRule{},
		TrimmedLengthRule{},
		OddPurchaseDateRule{},
		TimeOfPurchaseBetween2and4Rule{},
	}

	return &ReceiptPointsCalculator{Rules: rules}
}

func (r *ReceiptPointsCalculator) GenerateReceiptPoints(receipt models.Receipt) int {
	receiptPoints := 0

	for _, rule := range r.Rules {
		receiptPoints += rule.CalculatePoints(receipt)
	}

	return receiptPoints
}
