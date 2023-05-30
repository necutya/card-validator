package cardservice

import (
	"context"

	"github.com/necutya/card_validator/internal/models/domain"
	"github.com/necutya/card_validator/internal/models/dto"
)

type Service struct{}

func New() Service {
	return Service{}
}

func (s Service) ValidateCard(_ context.Context, cardDTO dto.ValidateCardRequestBody) error {
	card := domain.NewCardFromDto(cardDTO)
	return card.Validate()
}
