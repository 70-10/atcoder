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
		{s: "2019/04/30", expect: "Heisei"},
		{s: "2019/11/01", expect: "TBD"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.s))
	}
}
