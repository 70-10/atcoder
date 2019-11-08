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
		{a: 30, b: 100, expect: 100},
		{a: 12, b: 100, expect: 50},
		{a: 0, b: 100, expect: 0},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.b))
	}
}
