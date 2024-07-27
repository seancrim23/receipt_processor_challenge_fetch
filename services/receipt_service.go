package services

import "receipt_processor/models"

type ReceiptService interface {
	ProcessReceipt(models.Receipt) (string, error)
	GetReceiptPoints(string) (int, error)
}
