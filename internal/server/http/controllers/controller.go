package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/necutya/card_validator/internal/models/dto"
	log "github.com/sirupsen/logrus"
)

type CardService interface {
	ValidateCard(context.Context, dto.ValidateCardRequestBody) error
}

type Controller struct {
	cardService CardService
}

func New(service CardService) Controller {
	return Controller{
		cardService: service,
	}
}

func (*Controller) sendResponse(w http.ResponseWriter, r *http.Request, statusCode int, resp any) {
	log.Infof(
		"resp %s: %s - %d - %v",
		r.Method,
		r.RequestURI,
		statusCode,
		resp,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if resp != nil {
		respBody, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if resp != nil {
			if _, err = w.Write(respBody); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
