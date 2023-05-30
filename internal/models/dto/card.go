package dto

type ValidateCardRequestBody struct {
	CardNumber      string `json:"card_number"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
}

type ValidateCardResponseBody struct {
	Valid bool                   `json:"valid"`
	Error *CodeErrorResponseBody `json:"error,omitempty"`
}
