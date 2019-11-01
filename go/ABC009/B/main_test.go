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
		{nums: []int{100, 200, 200, 300}, expect: 200},
		{nums: []int{50, 370, 819, 433, 120}, expect: 433},
		{nums: []int{100, 100, 100, 200, 200, 200}, expect: 100},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.nums))
	}
}
