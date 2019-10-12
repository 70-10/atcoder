package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		expect int
	}{
		{s: "CSS", t: "CSR", expect: 2},
		{s: "SSR", t: "SSR", expect: 3},
		{s: "RRR", t: "SSS", expect: 0},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.s, test.t))
	}
}
