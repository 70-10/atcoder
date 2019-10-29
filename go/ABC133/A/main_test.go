package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		n      int
		a      int
		b      int
		expect int
	}{
		{n: 4, a: 2, b: 9, expect: 8},
		{n: 4, a: 2, b: 7, expect: 7},
		{n: 4, a: 2, b: 8, expect: 8},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.n, test.a, test.b))
	}
}
