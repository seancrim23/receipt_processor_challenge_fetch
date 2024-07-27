package services

import (
	"fmt"
	"log"
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

func TestProcessReceipt(t *testing.T) {
	service, _ := NewInMemoryReceiptService()

	t.Run("can successfully process a receipt", func(t *testing.T) {
		receiptId, err := service.ProcessReceipt(testReceipt1)
		if err != nil {
			fmt.Println("failure processing a receipt...")
			log.Fatal(err)
		}
		if len(receiptId) == 0 {
			t.Fatalf("no id created when receipt was created")
		}
	})
}

func TestGetReceiptPoints(t *testing.T) {
	service, _ := NewInMemoryReceiptService()

	t.Run("can successfully retrieve correct points amount", func(t *testing.T) {
		expectedPoints := 28
		service.receiptMap[testReceipt1.Id] = testReceipt1
		pointsResponse, err := service.GetReceiptPoints(testReceipt1.Id)
		if err != nil {
			fmt.Println("failure processing a receipt...")
			log.Fatal(err)
		}
		if expectedPoints != pointsResponse {
			t.Fatalf("expected %d points in response got %d", expectedPoints, pointsResponse)
		}
	})
}
