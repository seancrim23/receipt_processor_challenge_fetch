package utils

import (
	"receipt_processor/models"
	"testing"
)

var testReceipt1 = models.Receipt{
	Id:           "abc123",
	Retailer:     "Target",
	PurchaseDate: "2022-01-01",
	PurchaseTime: "13:01",
	Items: []models.Item{
		{
			ShortDescription: "Mountain Dew 12PK",
			Price:            "6.49",
		},
		{
			ShortDescription: "Emils Cheese Pizza",
			Price:            "12.25",
		},
		{
			ShortDescription: "Knorr Creamy Chicken",
			Price:            "1.26",
		},
		{
			ShortDescription: "Doritos Nacho Cheese",
			Price:            "3.35",
		},
		{
			ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
			Price:            "12.00",
		},
	},
	Total: "35.35",
}

var testReceipt2 = models.Receipt{
	Id:           "123abc",
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "14:33",
	Items: []models.Item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "9.00",
}

func TestAlphaNumericRule(t *testing.T) {
	t.Run("can calculate one point for every alphanumeric character in the retailer name", func(t *testing.T) {
		alphanumericRule := AlphanumericRule{}
		expectedPoints := 14

		calculatedPoints := alphanumericRule.CalculatePoints(testReceipt2)

		if expectedPoints != calculatedPoints {
			t.Fatalf("expected %d points from calculation but got %d", expectedPoints, calculatedPoints)
		}
	})
}

func TestRoundDollarTotalRule(t *testing.T) {
	t.Run("can return 50 points if the total is a round dollar amount with no cents", func(t *testing.T) {
		roundDollarTotalRule := RoundDollarTotalRule{}
		expectedPoints := 50

		calculatedPoints := roundDollarTotalRule.CalculatePoints(testReceipt2)

		if expectedPoints != calculatedPoints {
			t.Fatalf("expected %d points from calculation but got %d", expectedPoints, calculatedPoints)
		}
	})
}

func TestMultipleOfQuarterTotalRule(t *testing.T) {
	t.Run("can return 25 points if the total is a multiple of 0.25", func(t *testing.T) {
		multipleOfQuarterTotalRule := MultipleOfQuarterTotalRule{}
		expectedPoints := 25

		calculatedPoints := multipleOfQuarterTotalRule.CalculatePoints(testReceipt2)

		if expectedPoints != calculatedPoints {
			t.Fatalf("expected %d points from calculation but got %d", expectedPoints, calculatedPoints)
		}
	})
}

func TestTwoItemsRule(t *testing.T) {
	t.Run("can return 5 points for every two items on the receipt", func(t *testing.T) {
		twoItemsRule := TwoItemsRule{}
		expectedPoints := 10

		calculatedPoints := twoItemsRule.CalculatePoints(testReceipt1)

		if expectedPoints != calculatedPoints {
			t.Fatalf("expected %d points from calculation but got %d", expectedPoints, calculatedPoints)
		}
	})
}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func TestTrimmedLengthRule(t *testing.T) {
	t.Run("can return the correct number of points if trimmed length of item description is multiple of 3", func(t *testing.T) {
		trimmedLengthRule := TrimmedLengthRule{}
		expectedPoints := 6

		calculatedPoints := trimmedLengthRule.CalculatePoints(testReceipt1)

		if expectedPoints != calculatedPoints {
			t.Fatalf("expected %d points from calculation but got %d", expectedPoints, calculatedPoints)
		}
	})
}

func TestOddPurchaseDateRule(t *testing.T) {
	t.Run("can return 6 points if the day in the purchase date is odd", func(t *testing.T) {
		oddPurchaseDateRule := OddPurchaseDateRule{}
		expectedPoints := 6

		calculatedPoints := oddPurchaseDateRule.CalculatePoints(testReceipt1)

		if expectedPoints != calculatedPoints {
			t.Fatalf("expected %d points from calculation but got %d", expectedPoints, calculatedPoints)
		}
	})
}

func TestTimeOfPurchaseBetween2and4Rule(t *testing.T) {
	t.Run("can return 10 points if the time of purchase is after 2:00pm and before 4:00pm", func(t *testing.T) {
		timeOfPurchaseBetween2and4Rule := TimeOfPurchaseBetween2and4Rule{}
		expectedPoints := 10

		calculatedPoints := timeOfPurchaseBetween2and4Rule.CalculatePoints(testReceipt2)

		if expectedPoints != calculatedPoints {
			t.Fatalf("expected %d points from calculation but got %d", expectedPoints, calculatedPoints)
		}
	})
}
