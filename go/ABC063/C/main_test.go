package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{nums: []int{5, 10, 15}, expect: 25},
		{nums: []int{10, 10, 15}, expect: 35},
		{nums: []int{10, 20, 30}, expect: 0},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.nums))
	}
}
