package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		s      string
		expect int
	}{
		{s: "aabbbbaaca", expect: 5},
		{s: "aaaaa", expect: 1},
		{s: "xxzaffeeeeddfkkkkllq", expect: 10},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.s))
	}
}
