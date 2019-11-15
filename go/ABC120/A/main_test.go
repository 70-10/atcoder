package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		c      int
		expect int
	}{
		{a: 2, b: 11, c: 4, expect: 4},
		{a: 3, b: 9, c: 5, expect: 3},
		{a: 100, b: 1, c: 10, expect: 0},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.b, test.c))
	}
}
