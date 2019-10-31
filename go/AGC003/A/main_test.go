package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		s      string
		expect string
	}{
		{s: "SENW", expect: "Yes"},
		{s: "NSNNSNSN", expect: "Yes"},
		{s: "NSNNSNSNE", expect: "No"},
		{s: "NNEW", expect: "No"},
		{s: "W", expect: "No"},
		{s: "WE", expect: "Yes"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.s))
	}
}
