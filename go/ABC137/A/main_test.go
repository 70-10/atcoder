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
		{a: -13, b: 3, expect: -10},
		{a: 1, b: -33, expect: 34},
		{a: 13, b: 3, expect: 39},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.b))
	}
}
