package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		a      int
		p      int
		expect int
	}{
		{a: 1, p: 3, expect: 3},
		{a: 0, p: 1, expect: 0},
		{a: 32, p: 21, expect: 58},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.p))
	}
}
