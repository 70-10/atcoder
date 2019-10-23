package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		k      int
		x      int
		expect []int
	}{
		{k: 3, x: 7, expect: []int{5, 6, 7, 8, 9}},
		{k: 4, x: 0, expect: []int{-3, -2, -1, 0, 1, 2, 3}},
		{k: 1, x: 100, expect: []int{100}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.k, test.x))
	}
}
