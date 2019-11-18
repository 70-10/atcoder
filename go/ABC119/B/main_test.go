package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		o      []Otoshidama
		expect float64
	}{
		{o: []Otoshidama{
			{Value: 10000, Unit: "JPY"},
			{Value: 0.10000000, Unit: "BTC"},
		},
			expect: 48000.0,
		},
		{o: []Otoshidama{
			{Value: 100000000, Unit: "JPY"},
			{Value: 100.00000000, Unit: "BTC"},
			{Value: 0.00000001, Unit: "BTC"},
		},
			expect: 138000000.0038,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.o))
	}
}
