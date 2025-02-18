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
	"gopkg.in/validator.v2"
)

type ReceiptServer struct {
	service services.ReceiptService
	http.Handler
}

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
		utils.RespondWithError(w, code, errors.New(utils.INVALID_RECEIPT).Error())
		return
	}

	err = json.Unmarshal(reqBody, &receipt)
	if err != nil {
		code = 400
		fmt.Println(err)
		utils.RespondWithError(w, code, errors.New(utils.INVALID_RECEIPT).Error())
		return
	}
	if err = validator.Validate(receipt); err != nil {
		code = 400
		fmt.Println(err)
		utils.RespondWithError(w, code, errors.New(utils.INVALID_RECEIPT).Error())
		return
	}

	newReceiptId, err := h.service.ProcessReceipt(receipt)
	if err != nil {
		code = 400
		fmt.Println(err)
		utils.RespondWithError(w, code, errors.New(utils.INVALID_RECEIPT).Error())
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
	if err != nil {
		if err == errors.New(utils.NO_RECEIPT) {
			code = 404
			utils.RespondWithError(w, code, errors.New(utils.NO_RECEIPT).Error())
			return
		}
		code = 500
		utils.RespondWithError(w, code, err.Error())
		return
	}

	utils.RespondWithJSON(w, code, map[string]int{"points": points})
}
