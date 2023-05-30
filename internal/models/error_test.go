package models

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewCodeError(t *testing.T) {
	var (
		invalidCardNumberCodeError = CodeError{
			Code: InvalidCardNumberErrorCode,
			Err:  errors.New("invalid card number"),
		}
		expiredCardCodeError = CodeError{
			Code: ExpiredCardErrorCode,
			Err:  errors.New("expired date"),
		}
		baseCodeError = CodeError{
			Code: Code("001"),
			Err:  errors.New("system error"),
		}
	)

	type args struct {
		code Code
	}
	tests := []struct {
		name string
		args args
		want CodeError
	}{
		{
			name: "invalid card number case",
			args: args{
				code: InvalidCardNumberErrorCode,
			},
			want: invalidCardNumberCodeError,
		},
		{
			name: "expired card case",
			args: args{
				code: ExpiredCardErrorCode,
			},
			want: expiredCardCodeError,
		},
		{
			name: "base error case",
			args: args{
				code: Code("000"),
			},
			want: baseCodeError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCodeError(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCodeErrorByCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
