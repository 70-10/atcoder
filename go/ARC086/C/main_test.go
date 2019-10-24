package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		expect int
	}{
		{nums: []int{1, 1, 2, 2, 5}, k: 2, expect: 1},
		{nums: []int{1, 1, 2, 2}, k: 4, expect: 0},
		{nums: []int{5, 1, 3, 2, 4, 1, 1, 2, 3, 4}, k: 3, expect: 3},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.nums, test.k))
	}
}
