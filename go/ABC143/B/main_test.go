package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		a      []int
		expect int
	}{
		{a: []int{3, 1, 2}, expect: 11},
		{a: []int{5, 0, 7, 8, 3, 3, 2}, expect: 312},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a))
	}
}
