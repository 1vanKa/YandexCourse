package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPositive(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {
	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}
	_, err = Add(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}
	_, err = Add(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestAddZero(t *testing.T) {
	_, err := Add(0, 1)
	if err == nil {
		t.Error("first arg zero - expected error not be nil")
	}
	_, err = Add(1, 0)
	if err == nil {
		t.Error("second arg zero - expected error not be nil")
	}
	_, err = Add(0, 0)
	if err == nil {
		t.Error("all args zero - expected error not be nil")
	}
}

func TestEstimateValueTableDriven(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Small",
			args: args{value: 2},
			want: "small",
		},
		{
			name: "Test Medium",
			args: args{value: 10},
			want: "medium",
		},
		{
			name: "Test Big",
			args: args{value: 100},
			want: "big",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, EstimateValue(tt.args.value))
		})
	}
}
