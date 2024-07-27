package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"receipt_processor/models"
	"receipt_processor/services"
	"receipt_processor/utils"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ReceiptServer struct {
	service services.ReceiptService
	http.Handler
}

/*
Endpoint: Process Receipts
Path: /receipts/process
Method: POST
Payload: Receipt JSON
Response: JSON containing an id for the receipt.

Endpoint: Get Points
Path: /receipts/{id}/points
Method: GET
Response: A JSON object containing the number of points awarded.
A simple Getter endpoint that looks up the receipt by the ID and returns an object specifying the points awarded.

Example Response:

{ "points": 32 }

*/

func NewReceiptServer(service services.ReceiptService) (*ReceiptServer, error) {
	h := new(ReceiptServer)

	h.service = service

	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", h.processReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", h.getReceiptPoints).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug:          false,
	}).Handler(r)

	h.Handler = handler

	return h, nil
}

func (h *ReceiptServer) processReceipt(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var code = 201
	var err error
	var receipt models.Receipt

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		code = 400
		fmt.Println(err)
		utils.RespondWithError(w, code, err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &receipt)
	if err != nil {
		code = 400
		fmt.Println(err)
		utils.RespondWithError(w, code, err.Error())
		return
	}

	newReceiptId, err := h.service.ProcessReceipt(receipt)
	if err != nil {
		code = 500
		fmt.Println(err)
		utils.RespondWithError(w, code, err.Error())
		return
	}

	utils.RespondWithJSON(w, code, map[string]string{"id": newReceiptId})
}

func (h *ReceiptServer) getReceiptPoints(w http.ResponseWriter, r *http.Request) {
	var code = 200
	var err error

	receiptId := mux.Vars(r)["id"]
	if receiptId == "" {
		code = 400
		utils.RespondWithError(w, code, errors.New("no id passed to request").Error())
		return
	}

	points, err := h.service.GetReceiptPoints(receiptId)
	//determine what type of error and change code and return according error message
	if err != nil {
		code = 500
		utils.RespondWithError(w, code, err.Error())
		return
	}

	//make this into json of {"points": points}
	utils.RespondWithJSON(w, code, map[string]int{"points": points})
}
