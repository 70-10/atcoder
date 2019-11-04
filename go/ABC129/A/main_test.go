package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		p      int
		q      int
		r      int
		expect int
	}{
		{p: 1, q: 3, r: 4, expect: 4},
		{p: 3, q: 2, r: 3, expect: 5},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.p, test.q, test.r))
	}
}
