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
		{a: 2, b: 5, expect: 10},
		{a: 10, b: 5, expect: -1},
		{a: 9, b: 9, expect: 81},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.b))
	}
}
