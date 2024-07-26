package services

import (
	"errors"
	"receipt_processor/models"
	"receipt_processor/utils"

	"github.com/google/uuid"
)

type InMemoryReceiptService struct {
	receiptMap map[string]models.Receipt
}

func NewInMemoryReceiptService() (*InMemoryReceiptService, error) {
	receiptMap := map[string]models.Receipt{}
	return &InMemoryReceiptService{receiptMap: receiptMap}, nil
}

func (i *InMemoryReceiptService) ProcessReceipt(receipt models.Receipt) (*models.Receipt, error) {
	newReceiptId := uuid.New()
	receipt.Id = newReceiptId.String()
	i.receiptMap[newReceiptId.String()] = receipt
	return &receipt, nil
}

// what to return?
func (i *InMemoryReceiptService) GetReceiptPoints(id string) (int, error) {
	receipt, ok := i.receiptMap[id]
	if !ok {
		return 0, errors.New("error getting receipt")
	}
	receiptPointsCalculator := utils.NewReceiptPointsCalculator()
	points := receiptPointsCalculator.GenerateReceiptPoints(receipt)
	return points, nil
}
