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
		{nums: []int{2, 3}, expect: 6},
		{nums: []int{2, 5, 10, 1000000000000000000, 1000000000000000000}, expect: 1000000000000000000},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.nums))
	}
}
