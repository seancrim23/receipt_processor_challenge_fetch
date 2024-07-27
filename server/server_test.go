package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"receipt_processor/models"
	"testing"

	"github.com/gorilla/mux"
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

/*
GetShowsFunc   func(showQueryFilters map[string]string) (*[]models.Show, error)
GetShowFunc    func(id string) (*models.Show, error)
*/
type StubReceiptService struct {
	ProcessReceiptFunc   func(models.Receipt) (string, error)
	GetReceiptPointsFunc func(string) (int, error)

	receiptMap map[string]models.Receipt
}

func NewStubReceiptService() *StubReceiptService {
	newReceiptMap := map[string]models.Receipt{
		testReceipt1.Id: testReceipt1,
	}

	return &StubReceiptService{receiptMap: newReceiptMap}
}

func (s *StubReceiptService) ProcessReceipt(receipt models.Receipt) (string, error) {
	s.receiptMap[receipt.Id] = receipt
	return s.ProcessReceiptFunc(receipt)
}

func TestProcessReceipt(t *testing.T) {
	t.Run("can successfully process a receipt", func(t *testing.T) {
		service := &StubReceiptService{
			receiptMap: map[string]models.Receipt{},
		}
		expectedReceipt := models.Receipt{
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

		service.ProcessReceiptFunc = func(receipt models.Receipt) (string, error) {
			return expectedReceipt.Id, nil
		}

		server, _ := NewReceiptServer(service)

		req := httptest.NewRequest(http.MethodPost, "/receipts/process", receiptToJSON(expectedReceipt))
		res := httptest.NewRecorder()

		server.processReceipt(res, req)

		assertStatus(t, res.Code, http.StatusCreated)

		responseId := responseToId(res.Body)

		if responseId != expectedReceipt.Id {
			t.Errorf("the process receipt api call response %+v was not what was expected %+v", responseId, expectedReceipt.Id)
		}

		if len(service.receiptMap) != 1 {
			t.Fatalf("expected 1 receipt added but got %d", len(service.receiptMap))
		}
	})
}

func (s *StubReceiptService) GetReceiptPoints(id string) (int, error) {
	return s.GetReceiptPointsFunc(id)
}

func TestGetReceiptPoints(t *testing.T) {
	t.Run("gets expected points from get call", func(t *testing.T) {
		service := NewStubReceiptService()
		expectedPoints := 28
		service.GetReceiptPointsFunc = func(id string) (int, error) {
			return expectedPoints, nil
		}

		server, _ := NewReceiptServer(service)

		req := httptest.NewRequest(http.MethodGet, "/receipts/"+testReceipt1.Id+"/points/", nil)
		req = mux.SetURLVars(req, map[string]string{"id": testReceipt1.Id})
		res := httptest.NewRecorder()

		server.getReceiptPoints(res, req)

		assertStatus(t, res.Code, http.StatusOK)

		pointsResponse := responseToPoints(res.Body)

		if expectedPoints != pointsResponse {
			t.Fatalf("expected points of %d in response but got %d", expectedPoints, pointsResponse)
		}
	})
}

func responseToId(responseBody *bytes.Buffer) string {
	var idResponse map[string]string
	resBody, _ := io.ReadAll(responseBody)
	_ = json.Unmarshal(resBody, &idResponse)
	return idResponse["id"]
}

func responseToPoints(responseBody *bytes.Buffer) int {
	var pointsResponse map[string]int
	resBody, _ := io.ReadAll(responseBody)
	_ = json.Unmarshal(resBody, &pointsResponse)
	return pointsResponse["points"]
}

func receiptToJSON(receipt models.Receipt) io.Reader {
	b, _ := json.Marshal(receipt)
	return bytes.NewReader(b)
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
