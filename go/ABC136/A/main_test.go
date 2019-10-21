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
		{a: 6, b: 4, c: 3, expect: 1},
		{a: 8, b: 3, c: 9, expect: 4},
		{a: 12, b: 3, c: 7, expect: 0},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.b, test.c))
	}
}
