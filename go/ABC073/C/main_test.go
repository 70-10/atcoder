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
		{nums: []int{6, 2, 6}, expect: 1},
		{nums: []int{2, 5, 5, 2}, expect: 0},
		{nums: []int{12, 22, 16, 22, 18, 12}, expect: 2},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.nums))
	}
}
