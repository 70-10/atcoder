package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		x      int
		a      int
		expect int
	}{
		{x: 3, a: 5, expect: 0},
		{x: 6, a: 6, expect: 10},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.x, test.a))
	}
}
