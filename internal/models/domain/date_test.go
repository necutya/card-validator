package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestDate_Before(t *testing.T) {
	type fields struct {
		t time.Time
	}
	type args struct {
		other Date
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "before case",
			fields: fields{
				t: time.Date(2023, 05, 29, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				other: Date{
					t: time.Date(2023, 05, 30, 0, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
		{
			name: "after case",
			fields: fields{
				t: time.Date(2023, 05, 31, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				other: Date{
					t: time.Date(2023, 05, 30, 0, 0, 0, 0, time.UTC),
				},
			},
			want: false,
		},
		{
			name: "equal case",
			fields: fields{
				t: time.Date(2023, 05, 30, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				other: Date{
					t: time.Date(2023, 05, 30, 0, 0, 0, 0, time.UTC),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Date{
				t: tt.fields.t,
			}
			if got := d.Before(tt.args.other); got != tt.want {
				t.Errorf("Before() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDate(t *testing.T) {
	type args struct {
		y int
		m time.Month
		d int
	}
	tests := []struct {
		name string
		args args
		want Date
	}{
		{
			name: "simple case",
			args: args{
				y: 2023,
				m: time.April,
				d: 30,
			},
			want: Date{
				t: time.Date(2023, 05, 30, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDate(tt.args.y, tt.args.m, tt.args.d); got.t.Equal(tt.want.t) {
				t.Errorf("NewDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDateFromTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want Date
	}{
		{
			name: "simple case",
			args: args{
				t: time.Date(2023, 05, 30, 0, 0, 0, 0, time.UTC),
			},
			want: Date{
				t: time.Date(2023, 05, 30, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDateFromTime(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDateFromTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
