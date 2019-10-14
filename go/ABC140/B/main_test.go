package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		c      []int
		expect int
	}{
		{a: []int{3, 1, 2}, b: []int{2, 5, 4}, c: []int{3, 6}, expect: 14},
		{a: []int{2, 3, 4, 1}, b: []int{13, 5, 8, 24}, c: []int{45, 9, 15}, expect: 74},
		{a: []int{1, 2}, b: []int{50, 50}, c: []int{50}, expect: 150},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.b, test.c))
	}
}
