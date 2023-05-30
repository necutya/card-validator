package domain

import (
	"reflect"
	"testing"

	"github.com/necutya/card_validator/internal/models"

	"github.com/necutya/card_validator/internal/models/dto"
)

func TestCardNumber_Validate(t *testing.T) {
	tests := []struct {
		name string
		cn   CardNumber
		want error
	}{
		{
			name: "valid card case",
			cn:   "4111111111111111",
			want: nil,
		},
		{
			name: "invalid card case 1",
			cn:   "1111111111111",
			want: models.NewCodeError(models.InvalidCardNumberErrorCode),
		},
		{
			name: "invalid card case 2",
			cn:   "1111111111111111",
			want: models.NewCodeError(models.InvalidCardNumberErrorCode),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cn.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Validate(t *testing.T) {
	type fields struct {
		Number         CardNumber
		ExpirationDate Date
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "valid card case",
			fields: fields{
				Number:         CardNumber("4111111111111111"),
				ExpirationDate: NewDate(2028, 12, 00),
			},
			want: nil,
		},
		{
			name: "invalid card number case",
			fields: fields{
				Number:         CardNumber("1111111111111"),
				ExpirationDate: NewDate(2028, 12, 00),
			},
			want: models.NewCodeError(models.InvalidCardNumberErrorCode),
		},
		{
			name: "invalid expired date case",
			fields: fields{
				Number:         CardNumber("4111111111111111"),
				ExpirationDate: NewDate(2000, 12, 00),
			},
			want: models.NewCodeError(models.ExpiredCardErrorCode),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				Number:         tt.fields.Number,
				ExpirationDate: tt.fields.ExpirationDate,
			}
			if got := c.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCardFromDto(t *testing.T) {
	type args struct {
		cardDTO dto.ValidateCardRequestBody
	}
	tests := []struct {
		name string
		args args
		want Card
	}{
		{
			name: "simple case",
			args: args{
				cardDTO: dto.ValidateCardRequestBody{
					CardNumber:      "4111111111111111",
					ExpirationYear:  2028,
					ExpirationMonth: 12,
				},
			},
			want: Card{
				Number:         CardNumber("4111111111111111"),
				ExpirationDate: NewDate(2028, 12, 00),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCardFromDto(tt.args.cardDTO); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCardFromDto() = %v, want %v", got, tt.want)
			}
		})
	}
}
