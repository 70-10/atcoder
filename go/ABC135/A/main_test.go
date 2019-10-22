package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		a           int
		b           int
		expect      int
		expectError error
	}{
		{a: 2, b: 16, expect: 9, expectError: nil},
		{a: 0, b: 3, expect: 0, expectError: fmt.Errorf("IMPOSSIBLE")},
		{a: 998244353, b: 99824435, expect: 549034394, expectError: nil},
	}

	for _, test := range tests {
		actual, err := logic(test.a, test.b)
		assert.Equal(t, test.expect, actual)
		assert.Equal(t, test.expectError, err)
	}
}
