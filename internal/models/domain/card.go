package domain

import (
	"regexp"
	"time"

	"github.com/necutya/card_validator/internal/models"

	"github.com/necutya/card_validator/internal/models/dto"
)

var (
	// Just a simple regex patter to check for a card validness.
	// Note: this regular expression only validates the format of the card number,
	// not its validity as a real card number issued by a specific provider.
	cardValidationRegex = regexp.MustCompile(`\b4\d{15}\b`)
)

type CardNumber string

func (cn CardNumber) Validate() error {
	if !cardValidationRegex.MatchString(string(cn)) {
		return models.NewCodeError(models.InvalidCardNumberErrorCode)
	}

	return nil
}

type Card struct {
	Number         CardNumber
	ExpirationDate Date
}

func NewCardFromDto(cardDTO dto.ValidateCardRequestBody) Card {
	return Card{
		Number:         CardNumber(cardDTO.CardNumber),
		ExpirationDate: NewDate(cardDTO.ExpirationYear, time.Month(cardDTO.ExpirationMonth), 0),
	}
}

func (c Card) Validate() error {
	if err := c.Number.Validate(); err != nil {
		return err
	}

	if c.ExpirationDate.Before(NewDateFromTime(time.Now())) {
		return models.NewCodeError(models.ExpiredCardErrorCode)
	}

	return nil
}
