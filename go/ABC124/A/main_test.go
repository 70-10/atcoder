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
		{a: 5, b: 3, expect: 9},
		{a: 3, b: 4, expect: 7},
		{a: 6, b: 6, expect: 12},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.b))
	}
}
