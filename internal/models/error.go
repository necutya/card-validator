package models

import (
	"errors"
	"fmt"

	"github.com/necutya/card_validator/internal/models/dto"
)

const (
	SystemErrorCode            Code = "001"
	BadRequestErrorCode        Code = "100"
	InvalidCardNumberErrorCode Code = "101"
	ExpiredCardErrorCode       Code = "102"
)

var (
	codeMessageErrorMap = map[Code]error{
		InvalidCardNumberErrorCode: errors.New("invalid card number"),
		BadRequestErrorCode:        errors.New("bad request"),
		ExpiredCardErrorCode:       errors.New("expired date"),
		SystemErrorCode:            errors.New("system error"),
	}
)

type Code string

func (c Code) String() string {
	return string(c)
}

type CodeError struct {
	Code Code
	Err  error
}

func NewCodeError(code Code) CodeError {
	err, ok := codeMessageErrorMap[code]
	if !ok {
		return CodeError{
			Code: SystemErrorCode,
			Err:  codeMessageErrorMap[SystemErrorCode],
		}
	}

	return CodeError{
		Code: code,
		Err:  err,
	}
}

func NewSystemCodeError(err error) CodeError {
	return CodeError{
		Code: SystemErrorCode,
		Err:  fmt.Errorf("%w: %w", codeMessageErrorMap[SystemErrorCode], err),
	}
}

func NewBadRequestCodeError(err error) CodeError {
	return CodeError{
		Code: BadRequestErrorCode,
		Err:  fmt.Errorf("%w: %w", codeMessageErrorMap[BadRequestErrorCode], err),
	}
}

func (ce CodeError) Error() string {
	return fmt.Sprintf("%s - %s", ce.Code, ce.Err.Error())
}

func (ce CodeError) String() string {
	return ce.Error()
}

func (ce CodeError) ToDTO() *dto.CodeErrorResponseBody {
	return &dto.CodeErrorResponseBody{
		Code:    ce.Code.String(),
		Message: ce.Err.Error(),
	}
}
