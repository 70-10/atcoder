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
		{nums: []int{10, 8, 8, 6}, expect: 3},
		{nums: []int{15, 15, 15}, expect: 1},
		{nums: []int{50, 30, 50, 100, 50, 80, 30}, expect: 4},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.nums))
	}
}
