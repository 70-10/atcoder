package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		t      int
		expect int
	}{
		{a: 3, b: 5, t: 7, expect: 10},
		{a: 3, b: 2, t: 9, expect: 6},
		{a: 20, b: 20, t: 19, expect: 0},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.b, test.t))
	}
}
