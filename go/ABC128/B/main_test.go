package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		r      []Restaurant
		expect []int
	}{
		{
			r: []Restaurant{
				{ID: 1, Name: "khabarovsk", Point: 20},
				{ID: 2, Name: "moscow", Point: 10},
				{ID: 3, Name: "kazan", Point: 50},
				{ID: 4, Name: "kazan", Point: 35},
				{ID: 5, Name: "moscow", Point: 60},
				{ID: 6, Name: "khabarovsk", Point: 40},
			},
			expect: []int{3, 4, 6, 1, 5, 2},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.r))
	}
}
