package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		a      int
		s      string
		expect string
	}{
		{a: 3200, s: "pink", expect: "pink"},
		{a: 3199, s: "pink", expect: "red"},
		{a: 4049, s: "red", expect: "red"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a, test.s))
	}
}
