package services

import "receipt_processor/models"

type ReceiptService interface {
	ProcessReceipt(models.Receipt) (*models.Receipt, error)
	GetReceiptPoints(string) (int, error)
}
