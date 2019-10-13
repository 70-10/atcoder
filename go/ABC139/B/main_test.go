package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		expect int
	}{
		{a: 4, b: 10, expect: 3},
		{a: 8, b: 9, expect: 2},
		{a: 8, b: 8, expect: 1},
		{a: 8, b: 8, expect: 1},
		{a: 4, b: 20, expect: 7},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.b))
	}
}
