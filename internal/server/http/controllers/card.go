package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/necutya/card_validator/internal/models"

	"github.com/necutya/card_validator/internal/models/dto"
)

func (c *Controller) ValidateCard(w http.ResponseWriter, r *http.Request) {
	var req dto.ValidateCardRequestBody

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.sendErrValidationResponse(w, r, http.StatusBadRequest, models.NewBadRequestCodeError(err))
		return
	}

	if err := c.cardService.ValidateCard(r.Context(), req); err != nil {
		c.sendErrValidationResponse(w, r, http.StatusNotAcceptable, err)
		return
	}

	c.sendResponse(w, r, http.StatusOK, dto.ValidateCardResponseBody{
		Valid: true,
	})
}

func (c *Controller) sendErrValidationResponse(w http.ResponseWriter, r *http.Request, code int, err error) {
	var codeErr models.CodeError
	switch {
	case errors.As(err, &codeErr):
		c.sendResponse(w, r, code, dto.ValidateCardResponseBody{
			Valid: false,
			Error: codeErr.ToDTO(),
		})
	default:
		c.sendResponse(w, r, http.StatusInternalServerError, dto.ValidateCardResponseBody{
			Valid: false,
			Error: models.NewSystemCodeError(err).ToDTO(),
		})
	}
}
