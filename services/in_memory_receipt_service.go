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

func (i *InMemoryReceiptService) ProcessReceipt(receipt models.Receipt) (string, error) {
	newReceiptId := uuid.New()
	receipt.Id = newReceiptId.String()
	i.receiptMap[newReceiptId.String()] = receipt
	return receipt.Id, nil
}

func (i *InMemoryReceiptService) GetReceiptPoints(id string) (int, error) {
	receipt, receiptExists := i.receiptMap[id]
	if !receiptExists {
		return 0, errors.New(utils.NO_RECEIPT)
	}
	points := utils.NewReceiptPointsCalculator().GenerateReceiptPoints(receipt)
	return points, nil
}
