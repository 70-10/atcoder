package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{input: 2, expect: 8},
		{input: 1, expect: 1},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.input))
	}
}
