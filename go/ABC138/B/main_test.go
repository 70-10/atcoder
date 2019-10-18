package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		nums   []int
		expect float64
	}{
		{nums: []int{10, 30}, expect: 7.5},
		{nums: []int{200, 200, 200}, expect: 66.66666666666667},
		{nums: []int{1000}, expect: 1000},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.nums))
	}
}
